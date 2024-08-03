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

func (service *CourseService) GetCourse(courseID uint) (*models.Course, error) {
	course, err := service.repo.GetWithRelations(service.repo.DB, courseID, "Holes", "Rounds")
	return course, err
}

func (service *CourseService) GetAllCourses() (*[]models.Course, error) {
	courses, err := service.repo.GetAllWithRelations(service.repo.DB, "Holes")
	return courses, err
}

func (service *CourseService) EditCourseName(courseName string, courseID uint) error {
	tx := service.repo.Begin()

	defer func() {
		if r := recover(); r != nil || tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	course, err := service.repo.GetById(tx, courseID)
	if err != nil {
		return err
	}

	course.Name = courseName
	if err = service.repo.Update(tx, course); err != nil {
		return err
	}

	return nil
}

func (service *CourseService) InsertCourse(courseName string) (*models.Course, error) {
	tx := service.repo.Begin()
	defer func() {
		if r := recover(); r != nil || tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	course := models.Course{Name: courseName}
	if err := service.repo.Create(tx, &course); err != nil {
		return nil, err
	}
	return &course, nil
}

func (service *CourseService) InsertHoleToCourse(courseID uint, par uint, nthHole uint) error {
	tx := service.repo.Begin()

	defer func() {
		if r := recover(); r != nil || tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	course, err := service.repo.GetWithRelations(tx, courseID, "holes")

	if err != nil {
		return err
	}

	hole := models.Hole{Par: par, NthHole: nthHole, CourseID: courseID}
	course.Holes = append(course.Holes, hole)

	if err = service.repo.Update(tx, course); err != nil {
		return err
	}

	return nil
}
