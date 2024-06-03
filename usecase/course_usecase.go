package usecase

import (
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/entity/dto"
	"github.com/kelompok-2/ilmu-padi/repository"
	"github.com/kelompok-2/ilmu-padi/shared/model"
)

// Initialize Struct Course Usecase
type courseUsecase struct {
	courseRepository repository.CourseRepository
}

// Initialize Interface Course Sender Usecase
type CourseUsecase interface {
	CreateCourse(courses entity.Course, user string) (entity.Course, error)
	GetAllCourses(page, size int, user string) ([]entity.Course, model.Paging, error)
	GetCourseByID(id dto.CourseIDDto, user string) (entity.Course, error)
	UpdateCourse(id dto.CourseIDDto, courses entity.Course, user string) (entity.Course, error)
	DeleteCourse(id dto.CourseIDDto, user string) error
}

// Construction to Access Course Usecase
func NewCourseUsecase(courseRepository repository.CourseRepository) CourseUsecase {
	return &courseUsecase{
		courseRepository: courseRepository,
	}
}

// Create Course
func (c *courseUsecase) CreateCourse(courses entity.Course, user string) (entity.Course, error) {
	course := entity.Course{Title: courses.Title, Description: courses.Description, Category: courses.Category, Video_URL: courses.Video_URL, Duration: courses.Duration, Instructor_Name: courses.Instructor_Name, Rating: courses.Rating}
	if err := c.courseRepository.Create(course); err != nil {
		return entity.Course{}, err
	}
	return course, nil
}

// Get All Courses
func (c *courseUsecase) GetAllCourses(page, size int, user string) ([]entity.Course, model.Paging, error) {
	return c.courseRepository.FindAll(page, size)
}

// Get Course By ID
func (c *courseUsecase) GetCourseByID(id dto.CourseIDDto, user string) (entity.Course, error) {
	return c.courseRepository.FindByID(id)
}

// Update Course
func (c *courseUsecase) UpdateCourse(id dto.CourseIDDto, courses entity.Course, user string) (entity.Course, error) {
	course, err := c.courseRepository.FindByID(id)
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

	if err := c.courseRepository.Update(course); err != nil {
		return entity.Course{}, err
	}

	return course, nil
}

// Delete Course
func (c *courseUsecase) DeleteCourse(id dto.CourseIDDto, user string) error {
	return c.courseRepository.Delete(id)
}
