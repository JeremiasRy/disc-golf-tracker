package services

import (
	"disc-golf-tracker/backend/pkg/models"
	"disc-golf-tracker/backend/pkg/repositories"
)

type CourseService struct {
	repo *repositories.CrudRepository[models.Course]
}

func NewCourseService(repository *repositories.CrudRepository[models.Course]) CourseService {
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

func (service *CourseService) CreateCourse(courseName string) (*models.Course, error) {
	course := models.Course{Name: courseName}
	if err := service.repo.Create(service.repo.DB, &course); err != nil {
		return nil, err
	}
	return &course, nil
}
