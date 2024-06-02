package routes

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

	// Route User Course Favourite
	PostUserCourseFavourite    = "/user/course/favourite"
	GetUserCourseFavouriteList = "/user/course/favourite/:user_id"

	// Route Authentitacion
	Register       = "/register"
	Login          = "/login"
	Logout         = "/logout"
	ForgotPassword = "/forgot"

	// Route User
	GetUserList = "/profile"

	// Route Payment
	PostPayment = "/payment"
)
