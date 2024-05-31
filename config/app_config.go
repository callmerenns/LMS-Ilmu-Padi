package config

const (
	// Route Group
	ApiGroup = "/api/v1"

	// Route Course
	PostCourse    = "/courses"
	GetCourseList = "/courses"
	GetCourse     = "/courses/:id"
	GetCourseName = "/courses/category/:category"
	PutCourse     = "/courses/:id"
	DelCourse     = "/courses/:id"

	// Route User
	GetUserList = "/profile"

	// Route Payment
	PostPayment = "/payment"
)
