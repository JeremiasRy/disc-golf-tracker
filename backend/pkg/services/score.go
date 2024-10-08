package services

import (
	"disc-golf-tracker/backend/pkg/models"
	"disc-golf-tracker/backend/pkg/repositories"
)

type ScoreService struct {
	repo *repositories.CrudRepository[models.Score]
}

func NewScoreService(repository *repositories.CrudRepository[models.Score]) ScoreService {
	return ScoreService{repo: repository}
}

func (service *ScoreService) CreateScore(holeId uint, scoreCardId uint) (*models.Score, error) {
	tx := service.repo.Begin()
	defer func() {
		if r := recover(); r != nil || tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	conditions := map[string]interface{}{
		"scorecard_id": scoreCardId,
		"hole_id":      holeId,
	}

	score, err := service.repo.SearchByColumns(tx, conditions)

	if err != nil {
		score := models.Score{HoleID: holeId, ScorecardID: scoreCardId, Strokes: 0, Penalties: 0}
		if err := service.repo.Create(tx, &score); err != nil {
			return nil, err
		}
		return &score, nil
	}

	return score, nil
}

func (service *ScoreService) UpdateScore(scoreID uint, newStrokes uint, newPenalties uint) (*models.Score, error) {
	tx := service.repo.Begin()
	defer func() {
		if r := recover(); r != nil || tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	score, err := service.repo.GetById(tx, scoreID)
	if err != nil {
		return nil, err
	}

	if newStrokes != 0 {
		score.Strokes = newStrokes
	}

	if newPenalties != 0 {
		score.Penalties = newPenalties
	}

	err = service.repo.Update(tx, score)

	if err != nil {
		return nil, err
	}

	return score, nil
}
