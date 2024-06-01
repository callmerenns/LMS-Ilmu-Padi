package usecase

import (
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/entity/dto"
	"github.com/kelompok-2/ilmu-padi/repository"
)

// Initialize Struct Course Usecase
type CourseUsecase struct {
	courseRepository *repository.CourseRepository
}

// Construction to Access Course Usecase
func NewCourseUsecase(courseRepository *repository.CourseRepository) *CourseUsecase {
	return &CourseUsecase{courseRepository: courseRepository}
}

// Create Course
func (u *CourseUsecase) CreateCourse(courses entity.Course) (entity.Course, error) {
	course := entity.Course{Title: courses.Title, Description: courses.Description, Category: courses.Category, Video_URL: courses.Video_URL, Duration: courses.Duration, Instructor_Name: courses.Instructor_Name, Rating: courses.Rating}
	if err := u.courseRepository.Create(course); err != nil {
		return entity.Course{}, err
	}
	return course, nil
}

// Get All Courses
func (u *CourseUsecase) GetAllCourses() ([]entity.Course, error) {
	return u.courseRepository.FindAll()
}

// Get Course By ID
func (u *CourseUsecase) GetCourseByID(id dto.CourseIDDto) (entity.Course, error) {
	return u.courseRepository.FindByID(id)
}

// Update Course
func (u *CourseUsecase) UpdateCourse(id dto.CourseIDDto, courses entity.Course) (entity.Course, error) {
	course, err := u.courseRepository.FindByID(id)
	if err != nil {
		return entity.Course{}, err
	}

	course.Title = courses.Title
	course.Description = courses.Description
	course.Category = courses.Category
	course.Video_URL = courses.Video_URL
	course.Duration = courses.Duration
	course.Instructor_Name = courses.Instructor_Name
	course.Rating = courses.Rating

	if err := u.courseRepository.Update(course); err != nil {
		return entity.Course{}, err
	}

	return course, nil
}

// Delete Course
func (u *CourseUsecase) DeleteCourse(id uint) error {
	return u.courseRepository.Delete(id)
}
