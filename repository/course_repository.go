package repository

import (
	"github.com/kelompok-2/ilmu-padi/entity"
	"gorm.io/gorm"
)

type courseRepo struct {
	db *gorm.DB
}

type ICourseRepo interface {
	FindUserCoursesByUserID(uid string) ([]entity.Course, error)
}

func (n *courseRepo) FindUserCoursesByUserID(uid string) ([]entity.Course, error) {
	var courses []entity.Course
	err := n.db.Where("user_id = ?", uid).Find(&courses).Error
	if err != nil {
		return []entity.Course{}, err
	}
	return courses, nil
}

func NewCourseRepo(db *gorm.DB) ICourseRepo {
	return &courseRepo{
		db: db,
	}
}
