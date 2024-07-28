package repository

import (
	"gorm.io/gorm"
)

type Repository[T any] struct {
	DB *gorm.DB
}

func NewRepository[T any](db *gorm.DB) Repository[T] {
	return Repository[T]{DB: db}
}

func (repo *Repository[T]) Create(model *T) error {
	return repo.DB.Create(model).Error
}

func (repo *Repository[T]) GetById(id uint) (*T, error) {
	var entity T
	if err := repo.DB.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *Repository[T]) GetAll() ([]T, error) {
	var models []T
	if err := repo.DB.Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func (repo *Repository[T]) Update(model T) error {
	return repo.DB.Save(&model).Error
}

func (repo *Repository[T]) DeleteByID(id uint) error {
	var model T
	return repo.DB.Delete(&model, id).Error
}

func (repo *Repository[T]) GetWithRelations(id uint, relations ...string) (*T, error) {
	var model T
	query := repo.DB
	for _, relation := range relations {
		query = query.Preload(relation)
	}
	if err := query.First(&model, id).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (repo *Repository[T]) GetAllWithRelations(relations ...string) ([]T, error) {
	var models []T
	query := repo.DB
	for _, relation := range relations {
		query = query.Preload(relation)
	}
	if err := query.Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}
