package service

import (
	"errors"
	"ticket-api/entity"
	"ticket-api/repository"
)

type CategoryService interface {
	GetAll() ([]entity.Category, error)
	Create(input entity.Category) (*entity.Category, error)
	Update(id uint, input entity.Category) (*entity.Category, error)
	Delete(id uint) error
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(r repository.CategoryRepository) CategoryService {
	return &categoryService{r}
}

func (s *categoryService) GetAll() ([]entity.Category, error) {
	return s.repo.FindAll()
}

func (s *categoryService) Create(input entity.Category) (*entity.Category, error) {
	if existing, _ := s.repo.FindByName(input.Name); existing != nil {
		return nil, errors.New("category already exists")
	}
	err := s.repo.Create(&input)
	return &input, err
}

func (s *categoryService) Update(id uint, input entity.Category) (*entity.Category, error) {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("category not found")
	}

	category.Name = input.Name
	category.Description = input.Description

	err = s.repo.Update(category)
	return category, err
}

func (s *categoryService) Delete(id uint) error {
	category, err := s.repo.FindByID(id)
	if err != nil {
		return errors.New("category not found")
	}
	return s.repo.Delete(category)
}
