package usecase

import (
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/repository"
)

type userCoursesFavouriteUsecase struct {
	userCoursesFavouriteRepository repository.IUserCoursesFavouriteRepository
}

type IUserCoursesFavouriteUsecase interface {
	AddOrRemoveToFavourite(ucf entity.UserCoursesFavourite) (error, string)
	FindAllByUserID(user_id uint) ([]entity.Course, error)
}

func NewUserCoursesFavouriteUsecase(userCoursesFavouriteRepository repository.IUserCoursesFavouriteRepository) IUserCoursesFavouriteUsecase {
	return &userCoursesFavouriteUsecase{
		userCoursesFavouriteRepository: userCoursesFavouriteRepository,
	}
}

func (uc *userCoursesFavouriteUsecase) AddOrRemoveToFavourite(ucf entity.UserCoursesFavourite) (error, string) {
	return uc.userCoursesFavouriteRepository.AddOrRemoveToFavourite(ucf)
}

func (uc *userCoursesFavouriteUsecase) FindAllByUserID(user_id uint) ([]entity.Course, error) {
	return uc.userCoursesFavouriteRepository.FindAllByUserID(user_id)
}
