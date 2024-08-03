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

func (repo *Repository[T]) Begin() *gorm.DB {
	return repo.DB.Begin()
}

func (repo *Repository[T]) Commit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (repo *Repository[T]) Rollback(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (repo *Repository[T]) Create(tx *gorm.DB, model *T) error {
	return tx.Create(model).Error
}

func (repo *Repository[T]) GetById(tx *gorm.DB, id uint) (*T, error) {
	var entity T
	if err := repo.DB.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *Repository[T]) GetAll(tx *gorm.DB) (*[]T, error) {
	var models []T
	if err := tx.Find(&models).Error; err != nil {
		return nil, err
	}
	return &models, nil
}

func (repo *Repository[T]) Update(tx *gorm.DB, model *T) error {
	return tx.Save(model).Error
}

func (repo *Repository[T]) DeleteByID(tx *gorm.DB, id uint) error {
	var model T
	return tx.Delete(&model, id).Error
}

func (repo *Repository[T]) GetWithRelations(tx *gorm.DB, id uint, relations ...string) (*T, error) {
	var model T
	for _, relation := range relations {
		tx = tx.Preload(relation)
	}
	if err := tx.First(&model, id).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (repo *Repository[T]) GetAllWithRelations(tx *gorm.DB, relations ...string) (*[]T, error) {
	var models []T
	for _, relation := range relations {
		tx = tx.Preload(relation)
	}
	if err := tx.Find(&models).Error; err != nil {
		return nil, err
	}
	return &models, nil
}
