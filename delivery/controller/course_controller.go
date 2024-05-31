package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelompok-2/ilmu-padi/usecase"
)

type CourseController struct {
	courseUsecase *usecase.CourseUsecase
}

func NewCourseController(courseUsecase *usecase.CourseUsecase) *CourseController {
	return &CourseController{courseUsecase: courseUsecase}
}

func (ctrl *CourseController) CreateCourse(c *gin.Context) {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	instructorID := c.MustGet("user_id").(string) // Assume instructor is logged in

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := ctrl.courseUsecase.CreateCourse(input.Title, input.Description, instructorID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"course": course})
}

func (ctrl *CourseController) GetAllCourses(c *gin.Context) {
	courses, err := ctrl.courseUsecase.GetAllCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"courses": courses})
}

func (ctrl *CourseController) GetCourseByID(c *gin.Context) {
	var input struct {
		ID uint `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := ctrl.courseUsecase.GetCourseByID(input.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"course": course})
}

func (ctrl *CourseController) UpdateCourse(c *gin.Context) {
	var input struct {
		ID          uint   `uri:"id" binding:"required"`
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindUri(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := ctrl.courseUsecase.UpdateCourse(input.ID, input.Title, input.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"course": course})
}

func (ctrl *CourseController) DeleteCourse(c *gin.Context) {
	var input struct {
		ID uint `uri:"id" binding:"required"`
	}

	if err := c.ShouldBindUri(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.courseUsecase.DeleteCourse(input.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}
