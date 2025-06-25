package service

import (
	"errors"
	"ticket-api/entity"
	"ticket-api/repository"
)

type TicketService interface {
	Purchase(userID uint, eventID uint, quantity int) (*entity.Ticket, error)
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
