package entity

type UserCoursesFavourite struct {
	ID       uint   `gorm:"primarykey"`
	UserID   string `json:"user_id"`
	CourseID string `json:"course_id"`
}
