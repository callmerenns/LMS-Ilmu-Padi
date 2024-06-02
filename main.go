package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kelompok-2/ilmu-padi/config"
	"github.com/kelompok-2/ilmu-padi/config/routes"
	"github.com/kelompok-2/ilmu-padi/delivery/controller"
	"github.com/kelompok-2/ilmu-padi/delivery/middleware"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/repository"
	"github.com/kelompok-2/ilmu-padi/shared/service"
	"github.com/kelompok-2/ilmu-padi/usecase"
	midtrans "github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func initDB() *gorm.DB {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config(initDB): %v", err)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connection established successfully")

	db.AutoMigrate(&entity.User{}, &entity.Course{}, &entity.Subscription{}, &entity.Payment{}, &entity.UserCoursesFavourite{})

	log.Println("Database migrations completed successfully")

	return db
}

func main() {
	// delivery.NewServer().Run()
	db := initDB()
	if db == nil {
		log.Fatal("Database connection is nil")
	}

	defer db.Close()

	// Initialize Midtrans Snap client
	snapClient := snap.Client{}
	snapClient.New("YOUR_MIDTRANS_SERVER_KEY", midtrans.Sandbox)

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config(main): %v", err)
	}

	mailService := service.NewMailService(cfg.SmtpConfig)
	log.Println("MailService initialized successfully")

	// Layer Auth
	authRepository := repository.NewAuthRepository(db)
	log.Println("AuthRepository initialized successfully")
	authUsecase := usecase.NewAuthUsecase(authRepository, mailService)
	log.Println("AuthUseCase initialized successfully")
	authController := controller.NewAuthController(authUsecase)
	log.Println("AuthController initialized successfully")

	// Layer User
	userRepository := repository.NewUserRepository(db)
	log.Println("UserRepository initialized successfully")
	userUsecase := usecase.NewUserUsecase(userRepository)
	log.Println("UserUseCase initialized successfully")
	userController := controller.NewUserController(userUsecase)
	log.Println("UserController initialized successfully")

	// Layer Course
	courseRepository := repository.NewCourseRepository(db)
	log.Println("CourseRepository initialized successfully")
	courseUsecase := usecase.NewCourseUsecase(courseRepository)
	log.Println("CourseUseCase initialized successfully")
	courseController := controller.NewCourseController(courseUsecase)
	log.Println("CourseController initialized successfully")

	// Layer User Course Favourite
	userCourseFavouriteRepository := repository.NewUserCoursesFavouriteRepository(db)
	log.Println("UserCourseFavouriteRepository initialized successfully")
	userCourseFavouriteUsecase := usecase.NewUserCoursesFavouriteUsecase(userCourseFavouriteRepository)
	log.Println("UserCourseFavouriteUseCase initialized successfully")
	userCoursesFavouriteController := controller.NewUserCoursesFavouriteController(userCourseFavouriteUsecase)
	log.Println("UserCourseFavouriteController initialized successfully")

	// Layer Payment
	paymentRepository := repository.NewPaymentRepository(db)
	log.Println("PaymentRepository initialized successfully")
	paymentUsecase := usecase.NewPaymentUsecase(&snapClient, paymentRepository)
	log.Println("PaymentUseCase initialized successfully")
	paymentController := controller.NewPaymentController(paymentUsecase)
	log.Println("PaymentController initialized successfully")

	r := gin.Default()

	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)
	r.POST("/logout", authController.Logout)
	r.POST("/forgot-password", authController.ForgotPassword)
	r.POST("/reset-password", authController.ResetPassword)

	auth := r.Group("/auth")
	auth.Use(middleware.JWTAuthMiddleware())
	{
		// Route Profile
		auth.GET("/profile", userController.GetList)
		auth.GET("/profile/:id", userController.GetProfileByID)
		// auth.GET("/profile/:email", middleware.RoleMiddleware("admin"), userController.GetProfileByID)
		// auth.GET("/profile/:subscription-status", middleware.RoleMiddleware("admin"), userController.GetProfileByID)
		// auth.GET("/profile/:course", middleware.RoleMiddleware("admin"), userController.GetProfileByID)

		// Route Course
		auth.POST("/courses", courseController.CreateCourse)
		auth.GET("/courses", courseController.GetAllCourses)
		auth.GET("/courses/:id", courseController.GetCourseByID)
		auth.PUT("/courses/:id", courseController.UpdateCourse)
		auth.DELETE("/courses/:id", courseController.DeleteCourse)

		// Route User Course Favourite
		auth.POST(routes.PostUserCourseFavourite, userCoursesFavouriteController.AddOrRemoveCourseFavourite)
		auth.GET(routes.GetUserCourseFavouriteList, userCoursesFavouriteController.GetUserFavouriteList)

		// Route Payment
		auth.POST("/payment", paymentController.CreatePayment)
	}

	r.Run(cfg.ApiPort)
}
