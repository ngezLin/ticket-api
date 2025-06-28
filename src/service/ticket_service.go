package service

import (
	"errors"
	"ticket-api/src/entity"
	"ticket-api/src/repository"
	"time"
)

type TicketService interface {
	Purchase(userID uint, eventID uint, quantity int) (*entity.Ticket, error)
	GetMyTickets(userID uint) ([]entity.Ticket, error)
	GetMyTicketByID(userID uint, ticketID uint) (*entity.Ticket, error)
	Cancel(userID, ticketID uint) (*entity.Ticket, error)
	GetSalesReport() ([]entity.SalesReport, error)
}

type ticketService struct {
	ticketRepo repository.TicketRepository
	eventRepo  repository.EventRepository
	userRepo   repository.UserRepository
}

func NewTicketService(tRepo repository.TicketRepository, eRepo repository.EventRepository, uRepo repository.UserRepository) TicketService {
	return &ticketService{
		ticketRepo: tRepo,
		eventRepo:  eRepo,
		userRepo:   uRepo,
	}
}

func (s *ticketService) Purchase(userID uint, eventID uint, quantity int) (*entity.Ticket, error) {
	if quantity <= 0 {
		return nil, errors.New("quantity must be at least 1")
	}

	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	event, err := s.eventRepo.FindByID(eventID)
	if err != nil {
		return nil, errors.New("event not found")
	}

	if event.Capacity < quantity {
		return nil, errors.New("not enough capacity")
	}

	if event.Status != "active" {
		return nil, errors.New("event is not available")
	}

	totalPrice := event.Price * float64(quantity)
	if user.Balance < totalPrice {
		return nil, errors.New("insufficient balance")
	}

	// Decrease balance & capacity
	user.Balance -= totalPrice
	event.Capacity -= quantity
	s.userRepo.Update(user)
	s.eventRepo.Update(event)

	ticket := &entity.Ticket{
		UserID:     userID,
		EventID:    eventID,
		Quantity:   quantity,
		Status:     "active",
		TotalPrice: totalPrice,
	}

	err = s.ticketRepo.Create(ticket)
	if err != nil {
		return nil, err
	}

	ticket, err = s.ticketRepo.FindByID(ticket.ID)
	if err != nil {
		return nil, err
	}

	return ticket, err
}

func (s *ticketService) GetMyTickets(userID uint) ([]entity.Ticket, error) {
	return s.ticketRepo.FindByUserID(userID)
}

func (s *ticketService) GetMyTicketByID(userID, ticketID uint) (*entity.Ticket, error) {
	ticket, err := s.ticketRepo.FindByID(ticketID)
	if err != nil {
        return nil, err
    }
    if ticket.UserID != userID {
        return nil, errors.New("forbidden")
    }
    return ticket, nil
}

func (s *ticketService) Cancel(userID, ticketID uint) (*entity.Ticket, error) {
    ticket, err := s.ticketRepo.FindByID(ticketID)
    if err != nil {
        return nil, err
    }
    if ticket.UserID != userID {
        return nil, errors.New("forbidden")
    }
    if ticket.Status == "cancelled" {
        return nil, errors.New("ticket already cancelled")
    }

    event, _ := s.eventRepo.FindByID(ticket.EventID)
    if time.Now().After(event.StartTime) {
        return nil, errors.New("event already started – cannot cancel")
    }

    // update ticket & event capacity
    ticket.Status = "cancelled"
    event.Capacity += ticket.Quantity
    s.ticketRepo.Update(ticket)
    s.eventRepo.Update(event)

    // (opsional) refund:
    // user, _ := s.userRepo.FindByID(userID)
    // user.Balance += ticket.TotalPrice
    // s.userRepo.Update(user)

    return ticket, nil
}

func (s *ticketService) GetSalesReport() ([]entity.SalesReport, error) {
	return s.ticketRepo.GetSalesReport()
}