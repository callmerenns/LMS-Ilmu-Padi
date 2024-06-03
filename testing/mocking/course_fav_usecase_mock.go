package mocking

import (
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/stretchr/testify/mock"
)

type CourseFavouriteUsecaseMock struct {
	mock.Mock
}

// FindAllByUserID implements usecase.IUserCoursesFavouriteUsecase.
func (m *CourseFavouriteUsecaseMock) FindAllByUserID(user_id uint) ([]entity.Course, error) {
	panic("unimplemented")
}

func (m *CourseFavouriteUsecaseMock) AddOrRemoveToFavourite(userCourseFavourite entity.UserCoursesFavourite) (error, string) {
	args := m.Called(userCourseFavourite)
	return args.Error(0), args.String(1)
}
