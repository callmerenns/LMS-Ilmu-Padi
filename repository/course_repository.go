package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kelompok-2/ilmu-padi/entity"
)

type CourseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

func (r *CourseRepository) Create(course entity.Course) error {
	return r.db.Create(&course).Error
}

func (r *CourseRepository) FindAll() ([]entity.Course, error) {
	var courses []entity.Course
	if err := r.db.Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *CourseRepository) FindByID(id uint) (entity.Course, error) {
	var course entity.Course
	if err := r.db.First(&course, id).Error; err != nil {
		return course, err
	}
	return course, nil
}

func (r *CourseRepository) Update(course entity.Course) error {
	return r.db.Save(&course).Error
}

func (r *CourseRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Course{}, id).Error
}
