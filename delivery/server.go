package delivery

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kelompok-2/ilmu-padi/config"
	ctrl "github.com/kelompok-2/ilmu-padi/delivery/controller"
	rp "github.com/kelompok-2/ilmu-padi/repository"
	uc "github.com/kelompok-2/ilmu-padi/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// mandatory: usecase & service
type Server struct {
	engine  *gin.Engine
	cfg     *config.Config
	usecase uc.IUserUsecase
}

func (s *Server) initRoute() {
	r := s.engine.Group("/api/v1")
	// m := middlewares.NewAuthMiddleware(s.jwtService, s.userUsecase)

	ctrl.NewUserCtrl(s.usecase).Routing(r)
}

func NewServer() *Server {
	var err error
	var c *config.Config
	c, err = config.NewConfig()
	if err != nil {
		log.Fatal("Unexpected error:", err)
	}

	var db *gorm.DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.DbHost, c.DbUser, c.DbPassword, c.DbName, c.DbPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Unexpected error:", err)
	}

	// Inject Usecase, Service, Repo
	courseRepo := rp.NewCourseRepo(db)
	subcribtionRepo := rp.NewSubscribtionRepo(db)
	userRepo := rp.NewUserRepo(db, courseRepo, subcribtionRepo)

	userUsecase := uc.NewUserUsecase(userRepo)

	// mandatory: usecase & service
	return &Server{
		usecase: userUsecase,
		engine:  gin.Default(),
		cfg:     c,
	}
}

func (s *Server) Run() {
	s.initRoute()
	s.engine.Run(":8080")
}
