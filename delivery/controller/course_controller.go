package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/entity/dto"
	"github.com/kelompok-2/ilmu-padi/shared/common"
	"github.com/kelompok-2/ilmu-padi/usecase"
)

// Initialize Struct Course Controller
type CourseController struct {
	courseUsecase *usecase.CourseUsecase
}

// Construction to Access Course Controller
func NewCourseController(courseUsecase *usecase.CourseUsecase) *CourseController {
	return &CourseController{courseUsecase: courseUsecase}
}

// Create Course
func (ctrl *CourseController) CreateCourse(c *gin.Context) {
	input := entity.Course{}
	// instructorID := c.MustGet("user_id").(string) // Assume instructor is logged in

	if err := c.ShouldBindJSON(&input); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Status Bad Request")
		return
	}

	course, err := ctrl.courseUsecase.CreateCourse(input)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, "Status Internal Server Error")
		return
	}

	common.SendSingleResponse(c, course, "Success")
}

// Get All Courses
func (ctrl *CourseController) GetAllCourses(c *gin.Context) {
	courses, err := ctrl.courseUsecase.GetAllCourses()
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, "Status Internal Server Error")
		return
	}
	common.SendSingleResponse(c, courses, "Success")
}

// Get Course By ID
func (ctrl *CourseController) GetCourseByID(c *gin.Context) {
	input := dto.CourseIDDto{}

	if err := c.ShouldBindUri(&input); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Status Bad Request")
		return
	}

	course, err := ctrl.courseUsecase.GetCourseByID(input)
	if err != nil {
		common.SendErrorResponse(c, http.StatusNotFound, "Status Not Found")
		return
	}
	common.SendSingleResponse(c, course, "Success")
}

// Update Course
func (ctrl *CourseController) UpdateCourse(c *gin.Context) {
	input := entity.Course{}
	ID := dto.CourseIDDto{}

	if err := c.ShouldBindUri(&input); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Status Bad Request")
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Status Bad Request")
		return
	}

	course, err := ctrl.courseUsecase.UpdateCourse(ID, input)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, "Status Internal Server Error")
		return
	}

	common.SendSingleResponse(c, course, "Success")
}

// Delete Course
func (ctrl *CourseController) DeleteCourse(c *gin.Context) {
	var input struct {
		ID uint `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&input); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Status Bad Request")
		return
	}

	if err := ctrl.courseUsecase.DeleteCourse(input.ID); err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, "Status Internal Server Error")
		return
	}

	common.SendSuccessResponse(c, http.StatusOK, "Course Delete Success")
}
