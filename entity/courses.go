package entity

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ContentURL  string `json:"content_url"`
	UserID      string `json:"author_id"`
	IsFree      bool   `json:"is_free"`
	ADS         []Ads  `json:"ads"`
}
