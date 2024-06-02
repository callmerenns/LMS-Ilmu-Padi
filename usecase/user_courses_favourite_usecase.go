package usecase

import (
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/repository"
)

type UserCoursesFavouriteUsecase struct {
	userCoursesFavouriteRepository *repository.UserCoursesFavouriteRepository
}

func NewUserCoursesFavouriteUsecase(userCoursesFavouriteRepository *repository.UserCoursesFavouriteRepository) *UserCoursesFavouriteUsecase {
	return &UserCoursesFavouriteUsecase{
		userCoursesFavouriteRepository: userCoursesFavouriteRepository,
	}
}

func (uc *UserCoursesFavouriteUsecase) AddOrRemoveToFavourite(ucf entity.UserCoursesFavourite) (error, string) {
	return uc.userCoursesFavouriteRepository.AddOrRemoveToFavourite(ucf)
}

func (uc *UserCoursesFavouriteUsecase) GetFavouriteList(user_id uint) ([]entity.Course, error) {
	return uc.userCoursesFavouriteRepository.FindAllByUserID(user_id)
}
