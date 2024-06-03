package mocking

import (
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/stretchr/testify/mock"
)

type CourseFavouriteRepoMock struct {
	mock.Mock
}

func (m *UserRepoMock) AddOrRemoveCourseFavourite(userCourseFavourite entity.UserCoursesFavourite) (error, string) {
	args := m.Called(userCourseFavourite)
	return args.Error(0), args.String(1)
}
