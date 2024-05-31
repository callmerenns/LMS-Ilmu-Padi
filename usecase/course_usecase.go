package usecase

import (
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/repository"
)

type CourseUsecase struct {
	courseRepository *repository.CourseRepository
}

func NewCourseUsecase(courseRepository *repository.CourseRepository) *CourseUsecase {
	return &CourseUsecase{courseRepository: courseRepository}
}

func (u *CourseUsecase) CreateCourse(title, description string, instructorID string) (entity.Course, error) {
	course := entity.Course{Title: title, Description: description, InstructorID: instructorID}
	if err := u.courseRepository.Create(course); err != nil {
		return entity.Course{}, err
	}
	return course, nil
}

func (u *CourseUsecase) GetAllCourses() ([]entity.Course, error) {
	return u.courseRepository.FindAll()
}

func (u *CourseUsecase) GetCourseByID(id uint) (entity.Course, error) {
	return u.courseRepository.FindByID(id)
}

func (u *CourseUsecase) UpdateCourse(id uint, title, description string) (entity.Course, error) {
	course, err := u.courseRepository.FindByID(id)
	if err != nil {
		return entity.Course{}, err
	}

	course.Title = title
	course.Description = description

	if err := u.courseRepository.Update(course); err != nil {
		return entity.Course{}, err
	}

	return course, nil
}

func (u *CourseUsecase) DeleteCourse(id uint) error {
	return u.courseRepository.Delete(id)
}
