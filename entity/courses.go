package entity

import "gorm.io/gorm"

// Initialize Struct Course
type Course struct {
	gorm.Model
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	Category        string  `json:"category"`
	Video_URL       string  `json:"video_url"`
	Duration        string  `json:"duration"`
	Instructor_Name string  `json:"name_instructor"`
	Rating          float32 `json:"rating"`
}
