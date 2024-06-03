package testing

import (
	"testing"

	"github.com/kelompok-2/ilmu-padi/testing/mocking"
	"github.com/kelompok-2/ilmu-padi/usecase"
	"github.com/stretchr/testify/suite"
)

type CourseFavouriteUseCaseTestSuite struct {
	suite.Suite
	arm *mocking.CourseFavouriteRepoMock
	auc usecase.UserCoursesFavouriteUsecase
}

func (s *CourseFavouriteUseCaseTestSuite) SetupTest() {
	s.arm = new(mocking.CourseFavouriteRepoMock)
	s.auc = usecase.NewUserCoursesFavouriteUsecase(s.arm)
}

func TestCourseFavouriteUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CourseFavouriteUseCaseTestSuite))
}

func (s *CourseFavouriteUseCaseTestSuite) TestAddOrRemoveCourseFavourite_Added() {
	s.arm.On("AddOrRemoveToFavourite", payload).Return(nil, "Add to Favourite Executed")
	result, err := s.auc.AddOrRemoveToFavourite(payload)

	s.arm.AssertExpectations(s.T())

	s.NoError(err)
	s.Equal("Add to Favourite Executed", result)
}

func (s *CourseFavouriteUseCaseTestSuite) TestAddOrRemoveCourseFavourite_Removed() {
	s.arm.On("AddOrRemoveToFavourite", payload).Return(nil, "Remove from Favourite Executed")
	result, err := s.auc.AddOrRemoveToFavourite(payload)

	s.arm.AssertExpectations(s.T())

	s.NoError(err)
	s.Equal("Remove from Favourite Executed", result)
}
