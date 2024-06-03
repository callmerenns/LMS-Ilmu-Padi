package testing

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kelompok-2/ilmu-padi/testing/mocking"
	"github.com/stretchr/testify/suite"
)

type CourseFavouriteControllerTestSuite struct {
	suite.Suite
	rg  *gin.RouterGroup
	aum *mocking.CourseFavouriteUsecaseMock
	// amm *mocking.AuthMiddlewareMock
}

func (s *CourseFavouriteControllerTestSuite) SetupTest() {
	// s.amm = new(mocking.AuthMiddlewareMock)
	s.aum = new(mocking.CourseFavouriteUsecaseMock)
	r := gin.Default()
	gin.SetMode(gin.TestMode)
	rg := r.Group("/api/v1")

	// rg.Use(s.amm.CheckToken("user"))
	s.rg = rg
}

func TestAuthorControllerTestSuite(t *testing.T) {
	suite.Run(t, new(CourseFavouriteControllerTestSuite))
}

func (s *CourseFavouriteControllerTestSuite) TestAddOrRemoveCourseFavourite_Added() {
	// s.amm.On("CheckToken", mock.Anything).Return(nil)
	s.aum.On("AddOrRemoveCourseFavourite", payload).Return(nil, "Add to Favourite Executed")

	// authorController := controller.NewUserCoursesFavouriteController(s.aum, s.rg, s.amm)
	// authorController.Routing()

}
