package services

import (
	"disc-golf-tracker/backend/pkg/models"
	"disc-golf-tracker/backend/pkg/repositories"
)

type HoleService struct {
	repo *repositories.CrudRepository[models.Hole]
}

func NewHoleService(repository *repositories.CrudRepository[models.Hole]) HoleService {
	return HoleService{repo: repository}
}

func (service *HoleService) CreateHole(nthHole uint, par uint, courseID uint) (*models.Hole, error) {
	hole := models.Hole{NthHole: nthHole, Par: par, CourseID: courseID}
	if err := service.repo.Create(service.repo.DB, &hole); err != nil {
		return nil, err
	}
	return &hole, nil
}

func (service *HoleService) GetHole(holeID uint) (*models.Hole, error) {
	return service.repo.GetWithRelations(service.repo.DB, holeID, "Scores")
}
