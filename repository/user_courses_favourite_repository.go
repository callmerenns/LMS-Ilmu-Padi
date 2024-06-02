package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/kelompok-2/ilmu-padi/entity"
)

type UserCoursesFavouriteRepository struct {
	db *gorm.DB
}

func NewUserCoursesFavouriteRepository(db *gorm.DB) *UserCoursesFavouriteRepository {
	return &UserCoursesFavouriteRepository{db: db}
}

func (r *UserCoursesFavouriteRepository) AddOrRemoveToFavourite(userCourseFavourite entity.UserCoursesFavourite) (error, string) {
	var ucf entity.UserCoursesFavourite
	// if its already exits delete, else create
	if err := r.db.Where("user_id = ? AND course_id = ?", userCourseFavourite.UserID, userCourseFavourite.CourseID).First(&ucf).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return r.db.Create(&userCourseFavourite).Error, "Add to Favourite Executed"
		}
		return err, ""
	}

	return r.db.Where("user_id = ? AND course_id = ?", userCourseFavourite.UserID, userCourseFavourite.CourseID).Delete(&entity.UserCoursesFavourite{}).Error, "Remove from Favourite Executed"
}

func (r *UserCoursesFavouriteRepository) FindAllByUserID(userid uint) ([]entity.Course, error) {
	var listRaw []entity.UserCoursesFavourite
	if err := r.db.Where("user_id = ?", userid).Find(&listRaw).Error; err != nil {
		return []entity.Course{}, err
	}

	var list []entity.Course
	for _, userCourseFavourite := range listRaw {
		var course entity.Course
		if err := r.db.Where("id = ?", userCourseFavourite.CourseID).First(&course).Error; err != nil {
			return []entity.Course{}, err
		}

		list = append(list, course)
	}

	return list, nil
}
