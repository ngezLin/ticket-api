package service

import (
	"errors"
	"ticket-api/entity"
	"ticket-api/repository"
	"time"
)

type EventService interface {
	GetAll() ([]entity.Event, error)
	Create(event entity.Event) (*entity.Event, error)
	Update(id uint, event entity.Event) (*entity.Event, error)
	Delete(id uint) error
}

type eventService struct {
	repo repository.EventRepository
}

func NewEventService(r repository.EventRepository) EventService {
	return &eventService{r}
}

func (s *eventService) GetAll() ([]entity.Event, error) {
	return s.repo.FindAll()
}

func (s *eventService) Create(input entity.Event) (*entity.Event, error) {
	if input.Capacity < 0 || input.Price < 0 {
		return nil, errors.New("capacity and price must be ≥ 0")
	}

	if existing, err := s.repo.FindByName(input.Name); err == nil && existing != nil {
		return nil, errors.New("event already exists")
	}

	err := s.repo.Create(&input)
	return &input, err
}

func (s *eventService) Update(id uint, input entity.Event) (*entity.Event, error) {
	event, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("event not found")
	}

	if time.Now().After(event.StartTime) {
		return nil, errors.New("event already started, can't be edited")
	}

	event.Name = input.Name
	event.Description = input.Description
	event.CategoryID = input.CategoryID
	event.Capacity = input.Capacity
	event.Price = input.Price
	event.Status = input.Status
	event.StartTime = input.StartTime

	err = s.repo.Update(event)
	return event, err
}

func (s *eventService) Delete(id uint) error {
	event, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("event not found")
	}
	return s.repo.Delete(event)
}
