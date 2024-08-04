package services

import (
	"disc-golf-tracker/backend/pkg/models"
	"disc-golf-tracker/backend/pkg/repositories"
)

type RoundService struct {
	repo *repositories.CrudRepository[models.Round]
}

func NewRoundService(repository *repositories.CrudRepository[models.Round]) RoundService {
	return RoundService{repo: repository}
}

func (service *RoundService) GetRound(roundID uint) (*models.Round, error) {
	return service.repo.GetWithRelations(service.repo.DB, roundID, "ScoreCards.Scores")
}

func (service *RoundService) CreateRound(courseID uint) (*models.Round, error) {
	round := models.Round{CourseID: courseID}
	if err := service.repo.Create(service.repo.DB, &round); err != nil {
		return nil, err
	}
	return &round, nil
}
