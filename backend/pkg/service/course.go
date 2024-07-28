package services

import (
	"disc-golf-tracker/backend/pkg/models"
	"disc-golf-tracker/backend/pkg/repository"
)

type CourseService struct {
	repo *repository.Repository[models.Course]
}

func NewCourseService(repository *repository.Repository[models.Course]) CourseService {
	return CourseService{repo: repository}
}

func (service *CourseService) InsertCourse(courseName string) (*models.Course, error) {
	course := models.Course{Name: courseName}
	if err := service.repo.Create(&course); err != nil {
		return nil, err
	}
	return &course, nil
}

func (service *CourseService) InsertHoleToCourse(courseID uint, par uint, nthHole uint) error {
	tx := service.repo.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	course, err := service.repo.GetWithRelations(tx, courseID, "holes")

	if err != nil {
		tx.Rollback()
		return err
	}

	hole := models.Hole{Par: par, NthHole: nthHole, CourseID: courseID}
	course.Holes = append(course.Holes, hole)

	if err = service.repo.Update(tx, course); err != nil {
		tx.Rollback()
		return err
	}

	if err = service.repo.Commit(tx); err != nil {
		return err
	}

	return nil
}
