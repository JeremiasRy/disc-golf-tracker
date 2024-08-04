package services

import (
	"disc-golf-tracker/backend/pkg/models"
	"disc-golf-tracker/backend/pkg/repositories"
)

type UserService struct {
	repo *repositories.CrudRepository[models.User]
}

func NewUserService(repository *repositories.CrudRepository[models.User]) UserService {
	return UserService{repo: repository}
}

func (service *UserService) CreateUser(userName string) (*models.User, error) {
	user := models.User{Name: userName}

	if err := service.repo.Create(service.repo.DB, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (service *UserService) EditUserName(newName string, id uint) error {
	tx := service.repo.Begin()
	defer func() {
		if r := recover(); r != nil || tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	user, err := service.repo.GetById(tx, id)

	if err != nil {
		return err
	}

	user.Name = newName
	err = service.repo.Update(tx, user)
	if err != nil {
		return err
	}

	return nil
}

func (service *UserService) GetUser(id uint) (*models.User, error) {
	user, err := service.repo.GetWithRelations(service.repo.DB, id, "ScoreCards.Scores")
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) GetAllUsers() (*[]models.User, error) {
	users, err := service.repo.GetAll(service.repo.DB)
	if err != nil {
		return nil, err
	}
	return users, nil
}
