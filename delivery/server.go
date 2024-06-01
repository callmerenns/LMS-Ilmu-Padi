package delivery

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/kelompok-2/ilmu-padi/config"
	"github.com/kelompok-2/ilmu-padi/entity"
	_ "github.com/lib/pq"
)

type Server struct {
	// expenseUc  usecase.CourseUsecase
	// userUc     usecase.UserUsecase
	// authUsc    usecase
	// jwtService service.JwtService
	engine *gin.Engine
	host   string
}

var (
	db  *gorm.DB
	err error
)

func (s *Server) initRoute() {
	// rg := s.engine.Group(config.ApiGroup)

	// controller.NewAuthController(s.authUsc, rg).Route()
	// controller.NewUserController(s.userUc, rg, authMid).Route()
	// controller.NewExpenseController(s.expenseUc, rg, authMid).Route()

}

func (s *Server) Run() {
	s.initRoute()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("server not running on host %s, becauce error %v", s.host, err.Error()))
	}
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)

	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&entity.User{}, &entity.Course{}, &entity.Subscription{})

	// jwtService := service.NewJwtService(cfg.TokenConfig)
	// courseRepo := repository.NewCourseRepository(db)
	// userRepo := repository.NewUserRepository(db)

	// courseUc := usecase.NewExpenseUseCase(courseRepo)
	// userUc := usecase.NewUserUseCase(userRepo)
	// authUc := usecase.NewAuthUseCase(userUc, jwtService)

	// userRepository := repository.NewUserRepository(db)
	// userUsecase := usecase.NewUserUsecase(userRepository)
	// userController := controller.NewUserController(userUsecase)

	// courseRepository := repository.NewCourseRepository(db)
	// courseUsecase := usecase.NewCourseUsecase(courseRepository)
	// courseController := controller.NewCourseController(courseUsecase)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		// courseUc: courseUc,
		// userController: userController,
		engine: engine,
		host:   host,
	}
}
