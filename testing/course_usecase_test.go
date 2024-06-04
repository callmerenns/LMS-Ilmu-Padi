package testing

import (
	"testing"

	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/entity/dto"
	"github.com/kelompok-2/ilmu-padi/testing/mocking"
	"github.com/kelompok-2/ilmu-padi/usecase"
	"github.com/stretchr/testify/suite"
)

type CourseUsecaseTestSuite struct {
	suite.Suite
	arm *mocking.CourseRepoMock
	auc usecase.CourseUsecase
}

func (s *CourseUsecaseTestSuite) SetupTest() {
	s.arm = new(mocking.CourseRepoMock)
	s.auc = usecase.NewCourseUsecase(s.arm)
}

func (s *CourseUsecaseTestSuite) TestGetAllCourse_Success() {
	s.arm.On("FindAll").Return([]entity.Course{}, nil)
	courses, _, err := s.auc.GetAllCourses(1, 10, "")
	s.Nil(err)

	s.arm.AssertExpectations(s.T())
	s.Equal([]entity.Course{}, courses)
}

func (s *CourseUsecaseTestSuite) TestGetCourseByID_Success() {
	s.arm.On("FindByID", 1).Return(entity.Course{}, nil)
	course, err := s.auc.GetCourseByID(1, "")
	s.Nil(err)

	s.arm.AssertExpectations(s.T())
	s.Equal(entity.Course{}, course)
}

func (s *CourseUsecaseTestSuite) TestCreateCourse_Success() {
	s.arm.On("Create", entity.Course{}).Return(nil)
	course, err := s.auc.CreateCourse(entity.Course{}, "")
	s.Nil(err)

	s.arm.AssertExpectations(s.T())
	s.Equal(entity.Course{}, course)
}

func (s *CourseUsecaseTestSuite) TestUpdateCourse_Success() {
	// Create a sample course to return when FindByID is called
	existingCourse := entity.Course{
		ID:              1,
		Title:           "Old Title",
		Description:     "Old Description",
		Category:        "Old Category",
		Video_URL:       "old_video_url",
		Duration:        10,
		Instructor_Name: "Old Instructor",
		Rating:          4.5,
	}
	updatedCourse := entity.Course{
		ID:              1,
		Title:           "New Title",
		Description:     "New Description",
		Category:        "New Category",
		Video_URL:       "new_video_url",
		Duration:        12,
		Instructor_Name: "New Instructor",
		Rating:          4.8,
	}

	// Mock the FindByID and Update methods
	s.arm.On("FindByID", dto.CourseIDDto{ID: 1}).Return(existingCourse, nil)
	s.arm.On("Update", updatedCourse).Return(nil)

	// Call the UpdateCourse method
	result, err := s.auc.UpdateCourse(1, updatedCourse, "")

	// Assert that there were no errors and that the result matches the updated course
	s.Nil(err)
	s.Equal(updatedCourse, result)

	// Assert that the expectations were met
	s.arm.AssertExpectations(s.T())
}

func (s *CourseUsecaseTestSuite) TestDeleteCourse_Success() {
	s.arm.On("Delete", 1).Return(nil)
	err := s.auc.DeleteCourse(1, "")
	s.Nil(err)

	s.arm.AssertExpectations(s.T())
}

func TestCourseUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(CourseUsecaseTestSuite))
}
