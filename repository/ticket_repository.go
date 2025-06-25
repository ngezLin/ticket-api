package repository

import (
	"ticket-api/entity"

	"gorm.io/gorm"
)

type TicketRepository interface {
	Create(ticket *entity.Ticket) error
	FindByID(id uint) (*entity.Ticket, error)
	FindByUserID(userID uint) ([]entity.Ticket, error)
	Update(ticket *entity.Ticket) error
}

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{db}
}

func (r *ticketRepository) Create(ticket *entity.Ticket) error {
	return r.db.Create(ticket).Error
}

func (r *ticketRepository) FindByID(id uint) (*entity.Ticket, error) {
	var ticket entity.Ticket
	err := r.db.
		Preload("Event.Category").Preload("User").First(&ticket, id).Error
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (r *ticketRepository) FindByUserID(userID uint) ([]entity.Ticket, error) {
	var tickets []entity.Ticket
	err := r.db.Preload("Event").Where("user_id = ?", userID).Find(&tickets).Error
	return tickets, err
}

func (r *ticketRepository) Update(ticket *entity.Ticket) error {
	return r.db.Save(ticket).Error
}