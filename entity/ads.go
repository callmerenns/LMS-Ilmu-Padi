package entity

import "gorm.io/gorm"

type Ads struct {
	gorm.Model
	ID       string `json:"id"`
	Content  string `json:"content"`
	CourseID string `json:"course_id"`
}
