package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kelompok-2/ilmu-padi/delivery/controller"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/testing/mocking"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CourseControllerTestSuite struct {
	suite.Suite
	rg  *gin.RouterGroup
	aum *mocking.CourseUsecaseMock
	amm *mocking.AuthMiddlewareMock
}

func (s *CourseControllerTestSuite) SetupTest() {
	s.amm = new(mocking.AuthMiddlewareMock)
	s.aum = new(mocking.CourseUsecaseMock)
	r := gin.Default()
	gin.SetMode(gin.TestMode)
	rg := r.Group("/api/v1")
	rg.Use(s.amm.RequireToken("user", "admin"))
	s.rg = rg
}

func (s *CourseControllerTestSuite) TestGetAllCourse_Success() {
	s.amm.On("RequireToken", "admin", "instructor", "user").Return(nil)
	s.aum.On("GetAllCourse").Return([]entity.Course{}, nil)
	coursesCtrl := controller.NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	req, err := http.NewRequest("GET", "/api/v1/asset/courses", nil)
	assert.NoError(s.T(), err)

	record := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(record)
	c.Request = req

	coursesCtrl.GetAllCourses(c)
	assert.Equal(s.T(), http.StatusOK, record.Code)
}

func (s *CourseControllerTestSuite) TestGetAllCourse_Failed() {
	s.amm.On("RequireToken", "admin", "instructor", "user").Return(nil)
	s.aum.On("GetAllCourse").Return([]entity.Course{}, fmt.Errorf("Get All Course Failed"))
	coursesCtrl := controller.NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	req, err := http.NewRequest("GET", "/api/v1/asset/courses", nil)
	assert.NoError(s.T(), err)

	record := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(record)
	c.Request = req

	coursesCtrl.GetAllCourses(c)
	assert.Equal(s.T(), http.StatusInternalServerError, record.Code)
}

func (s *CourseControllerTestSuite) TestGetCourse_Success() {
	s.amm.On("RequireToken", "admin", "instructor", "user").Return(nil)
	s.aum.On("GetCourse", "1").Return(entity.Course{}, nil)
	coursesCtrl := controller.NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	req, err := http.NewRequest("GET", "/api/v1/asset/course/1", nil)
	assert.NoError(s.T(), err)

	record := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(record)
	c.Request = req

	coursesCtrl.GetCourseByID(c)
	assert.Equal(s.T(), http.StatusOK, record.Code)
}

func (s *CourseControllerTestSuite) TestGetCourse_Failed() {
	s.amm.On("RequireToken", "admin", "instructor", "user").Return(nil)
	s.aum.On("GetCourse", "1").Return(entity.Course{}, nil)
	coursesCtrl := controller.NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	req, err := http.NewRequest("GET", "/api/v1/asset/course/1", nil)
	assert.NoError(s.T(), err)

	record := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(record)
	c.Request = req

	coursesCtrl.GetCourseByID(c)
	assert.Equal(s.T(), http.StatusOK, record.Code)
}

func (s *CourseControllerTestSuite) TestCreateCourse_Success() {
	payload := entity.Course{
		ID:              1,
		Title:           "Test",
		Description:     "Test",
		Category:        "Test",
		Video_URL:       "Test",
		Duration:        "0",
		Instructor_Name: "0",
		Rating:          0,
		UserId:          "0",
	}

	s.amm.On("RequireToken", "admin", "instructor").Return(nil)
	s.aum.On("CreateCourse", payload, "").Return(payload, nil)

	coursesCtrl := controller.NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	jsonPayload, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "/api/v1/asset/courses", bytes.NewBuffer(jsonPayload))
	assert.NoError(s.T(), err)
	req.Header.Set("Content-Type", "application/json")

	record := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(record)
	c.Request = req

	coursesCtrl.CreateCourse(c)

	assert.Equal(s.T(), http.StatusOK, record.Code)
}

func (s *CourseControllerTestSuite) TestCreateCourse_Failed() {
	s.amm.On("RequireToken", "admin", "instructor").Return(nil)
	s.aum.On("CreateCourse", entity.Course{}).Return(fmt.Errorf("Create Course Failed"))
	coursesCtrl := controller.NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	req, err := http.NewRequest("POST", "/api/v1/asset/courses", nil)
	assert.NoError(s.T(), err)

	record := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(record)
	c.Request = req

	coursesCtrl.CreateCourse(c)
	assert.Equal(s.T(), http.StatusInternalServerError, record.Code)
}

func (s *CourseControllerTestSuite) TestUpdateCourse_Success() {
	s.amm.On("RequireToken", "admin", "instructor").Return(nil)
	s.aum.On("UpdateCourse", entity.Course{}).Return(nil)
	coursesCtrl := controller.NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	req, err := http.NewRequest("PUT", "/api/v1/asset/courses/v1", nil)
	assert.NoError(s.T(), err)

	record := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(record)
	c.Request = req

	coursesCtrl.UpdateCourse(c)
	assert.Equal(s.T(), http.StatusOK, record.Code)
}

func (s *CourseControllerTestSuite) TestUpdateCourse_Failed() {
	s.amm.On("RequireToken", "admin", "instructor").Return(nil)
	s.aum.On("UpdateCourse", entity.Course{}).Return(fmt.Errorf("Update Course Failed"))
	coursesCtrl := controller.NewCourseController(s.aum, s.rg, s.amm)
	coursesCtrl.Route()

	req, err := http.NewRequest("PUT", "/api/v1/asset/courses/1", nil)
	assert.NoError(s.T(), err)

	record := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(record)
	c.Request = req

	coursesCtrl.UpdateCourse(c)
	assert.Equal(s.T(), http.StatusInternalServerError, record.Code)
}

func TestCourseControllerTestSuite(t *testing.T) {
	suite.Run(t, new(CourseControllerTestSuite))
}
