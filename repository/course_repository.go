package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/entity/dto"
)

// Initialize Struct Course Repository
type CourseRepository struct {
	db *gorm.DB
}

// Construction to Access Course Repository
func NewCourseRepository(db *gorm.DB) *CourseRepository {
	if db == nil {
		log.Fatal("Database connection is nil CourseRepository")
	}

	return &CourseRepository{db: db}
}

// Create
func (r *CourseRepository) Create(course entity.Course) error {
	if r.db == nil {
		log.Fatal("Database connection is nil in Create")
	}

	return r.db.Create(&course).Error
}

// Find All
func (r *CourseRepository) FindAll() ([]entity.Course, error) {
	if r.db == nil {
		log.Fatal("Database connection is nil in FindAll")
	}

	var courses []entity.Course
	if err := r.db.Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

// Find By ID
func (r *CourseRepository) FindByID(id dto.CourseIDDto) (entity.Course, error) {
	if r.db == nil {
		log.Fatal("Database connection is nil in FindByID")
	}

	var course entity.Course
	if err := r.db.First(&course, id).Error; err != nil {
		return course, err
	}
	return course, nil
}

// Update
func (r *CourseRepository) Update(course entity.Course) error {
	if r.db == nil {
		log.Fatal("Database connection is nil in Update")
	}

	return r.db.Save(&course).Error
}

// Delete
func (r *CourseRepository) Delete(id uint) error {
	if r.db == nil {
		log.Fatal("Database connection is nil in Delete")
	}

	return r.db.Delete(&entity.Course{}, id).Error
}
