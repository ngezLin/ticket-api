package repository

import (
	"ticket-api/src/entity"

	"gorm.io/gorm"
)

type EventRepository interface {
	FindAll() ([]entity.Event, error)
	FindByID(id uint) (*entity.Event, error)
	FindByName(name string) (*entity.Event, error)
	Create(event *entity.Event) error
	Update(event *entity.Event) error
	Delete(event *entity.Event) error
	FindAllActive() ([]entity.Event, error)
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) FindAll() ([]entity.Event, error) {
	var events []entity.Event
	err := r.db.Preload("Category").Find(&events).Error
	return events, err
}

func (r *eventRepository) FindByID(id uint) (*entity.Event, error) {
	var event entity.Event
	err := r.db.Preload("Category").First(&event, id).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *eventRepository) FindByName(name string) (*entity.Event, error) {
	var event entity.Event
	err := r.db.Where("name = ?", name).First(&event).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *eventRepository) Create(event *entity.Event) error {
	return r.db.Create(event).Error
}

func (r *eventRepository) Update(event *entity.Event) error {
	return r.db.Save(event).Error
}

func (r *eventRepository) Delete(event *entity.Event) error {
	return r.db.Delete(event).Error
}

func (r *eventRepository) FindAllActive() ([]entity.Event, error) {
	var events []entity.Event
	err := r.db.Preload("Category").Where("status = ?", "active").Find(&events).Error
	return events, err
}