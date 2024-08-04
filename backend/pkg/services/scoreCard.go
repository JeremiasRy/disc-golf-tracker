package services

import (
	"disc-golf-tracker/backend/pkg/models"
	"disc-golf-tracker/backend/pkg/repositories"
)

type ScoreCardService struct {
	repo *repositories.CrudRepository[models.ScoreCard]
}

func NewScoreCardService(repository *repositories.CrudRepository[models.ScoreCard]) ScoreCardService {
	return ScoreCardService{repo: repository}
}

func (service *ScoreCardService) CreateScoreCard(roundID uint, userID uint) (*models.ScoreCard, error) {
	scoreCard := models.ScoreCard{RoundID: roundID, UserID: userID}
	if err := service.repo.Create(service.repo.DB, &scoreCard); err != nil {
		return nil, err
	}
	return &scoreCard, nil
}
