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
	GetSalesReport() ([]entity.SalesReport, error)
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
	err := r.db.
		Preload("Event").
		Preload("Event.Category").
		Where("user_id = ?", userID).
		Find(&tickets).Error
	return tickets, err
}

func (r *ticketRepository) Update(ticket *entity.Ticket) error {
	return r.db.Save(ticket).Error
}

func (r *ticketRepository) GetSalesReport() ([]entity.SalesReport, error) {
	var reports []entity.SalesReport

		err := r.db.
		Table("tickets").
		Select("events.id as event_id, events.name as event_name, SUM(tickets.quantity) as tickets_sold, SUM(tickets.total_price) as total_revenue, events.capacity as remaining_quota, events.status").
		Joins("JOIN events ON events.id = tickets.event_id").
		Where("tickets.status = ?", "active").
		Group("events.id, events.name, events.capacity, events.status").
		Scan(&reports).Error

	return reports, err
}