package entity

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	ID            string          `json:"id"`
	Title         string          `json:"title"`
	Description   string          `json:"description"`
	InstructorID  string          `json:"instructor_id"`
	CourseContent []CourseContent `json:"course_content"`
}
