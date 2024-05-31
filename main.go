package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/kelompok-2/ilmu-padi/config"
	"github.com/kelompok-2/ilmu-padi/delivery/controller"
	"github.com/kelompok-2/ilmu-padi/delivery/middleware"
	"github.com/kelompok-2/ilmu-padi/email"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/repository"
	"github.com/kelompok-2/ilmu-padi/usecase"
	_ "github.com/lib/pq"
	midtrans "github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

var db *gorm.DB
var err error

func initDB() {
	cfg, _ := config.NewConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)

	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&entity.User{}, &entity.Course{}, &entity.Subscription{}, &entity.VerificationToken{}, &entity.CourseContent{})
}

func main() {
	// delivery.NewServer().Run()
	cfg, _ := config.NewConfig()
	// Initialize Midtrans Snap client
	snapClient := snap.Client{}
	snapClient.New("YOUR_MIDTRANS_SERVER_KEY", midtrans.Sandbox)

	userRepository := repository.NewUserRepository(db)
	emailSender := email.NewSMTPEmailSender("smtp.example.com", 587, "your-email@example.com", "your-email-password")
	userUsecase := usecase.NewUserUsecase(userRepository, emailSender)
	userController := controller.NewUserController(userUsecase)

	courseRepository := repository.NewCourseRepository(db)
	courseUsecase := usecase.NewCourseUsecase(courseRepository)
	courseController := controller.NewCourseController(courseUsecase)

	paymentRepository := repository.NewPaymentRepository(db)
	paymentUsecase := usecase.NewPaymentUsecase(&snapClient, paymentRepository)
	paymentController := controller.NewPaymentController(paymentUsecase)

	r := gin.Default()

	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)
	r.POST("/forgot-password", userController.ForgotPassword)
	r.POST("/reset-password", userController.ResetPassword)

	auth := r.Group("/auth")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		auth.GET("/profile", func(c *gin.Context) {
			userID := c.MustGet("user_id").(uint)
			user, err := userRepository.FindByID(userID)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"user": user})
		})

		auth.POST("/courses", middleware.RoleMiddleware("instructor"), courseController.CreateCourse)
		auth.GET("/courses", courseController.GetAllCourses)
		auth.GET("/courses/:id", courseController.GetCourseByID)
		auth.PUT("/courses/:id", middleware.RoleMiddleware("instructor"), courseController.UpdateCourse)
		auth.DELETE("/courses/:id", middleware.RoleMiddleware("admin"), courseController.DeleteCourse)

		auth.POST("/payment", paymentController.CreatePayment)
	}

	r.Run(cfg.ApiPort)
}
