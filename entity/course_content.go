package entity

import "gorm.io/gorm"

type CourseContent struct {
	gorm.Model
	ID          string `json:"id"`
	CourseID    string `json:"course_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	VideoURL    string `json:"video_url"`
}
