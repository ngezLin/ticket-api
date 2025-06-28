package repository

import (
	"errors"
	"ticket-api/src/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll() ([]entity.Category, error)
	FindByID(id uint) (*entity.Category, error)
	FindByName(name string) (*entity.Category, error)
	Create(category *entity.Category) error
	Update(category *entity.Category) error
	Delete(category *entity.Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) FindAll() ([]entity.Category, error) {
	var categories []entity.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) FindByID(id uint) (*entity.Category, error) {
	var category entity.Category
	err := r.db.First(&category, id).Error
	return &category, err
}

func (r *categoryRepository) FindByName(name string) (*entity.Category, error) {
	var category entity.Category
	err := r.db.Where("name = ?", name).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) Create(category *entity.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) Update(category *entity.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(category *entity.Category) error {
	return r.db.Delete(category).Error
}
