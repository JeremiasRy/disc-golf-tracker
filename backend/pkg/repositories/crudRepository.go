package repositories

import (
	"gorm.io/gorm"
)

type CrudRepository[T any] struct {
	DB *gorm.DB
}

func NewRepository[T any](db *gorm.DB) CrudRepository[T] {
	return CrudRepository[T]{DB: db}
}

func (repo *CrudRepository[T]) Begin() *gorm.DB {
	return repo.DB.Begin()
}

func (repo *CrudRepository[T]) Commit(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (repo *CrudRepository[T]) Rollback(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (repo *CrudRepository[T]) Create(tx *gorm.DB, model *T) error {
	return tx.Create(model).Error
}

func (repo *CrudRepository[T]) GetById(tx *gorm.DB, id uint) (*T, error) {
	var entity T
	if err := repo.DB.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *CrudRepository[T]) GetAll(tx *gorm.DB) (*[]T, error) {
	var models []T
	if err := tx.Find(&models).Error; err != nil {
		return nil, err
	}
	return &models, nil
}

func (repo *CrudRepository[T]) Update(tx *gorm.DB, model *T) error {
	return tx.Save(model).Error
}

func (repo *CrudRepository[T]) DeleteByID(tx *gorm.DB, id uint) error {
	var model T
	return tx.Delete(&model, id).Error
}

func (repo *CrudRepository[T]) GetWithRelations(tx *gorm.DB, id uint, relations ...string) (*T, error) {
	var model T
	for _, relation := range relations {
		tx = tx.Preload(relation)
	}
	if err := tx.First(&model, id).Error; err != nil {
		return nil, err
	}
	return &model, nil
}

func (repo *CrudRepository[T]) GetAllWithRelations(tx *gorm.DB, relations ...string) (*[]T, error) {
	var models []T
	for _, relation := range relations {
		tx = tx.Preload(relation)
	}
	if err := tx.Find(&models).Error; err != nil {
		return nil, err
	}
	return &models, nil
}

func (repo *CrudRepository[T]) SearchByColumns(tx *gorm.DB, conditions map[string]interface{}) (*T, error) {
	var model T
	query := tx
	for column, value := range conditions {
		query = query.Where(column+" = ?", value)
	}
	if err := query.First(&model).Error; err != nil {
		return nil, err
	}
	return &model, nil
}
