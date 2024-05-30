package delivery

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kelompok-2/ilmu-padi/config"
	rp "github.com/kelompok-2/ilmu-padi/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Note: All usecase / service should be injected
type Server struct {
	engine *gin.Engine
	cfg    *config.Config
}

func (s *Server) initRoute() {

}

func NewServer(cfg *config.Config) *Server {
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

	return &Server{
		engine: gin.Default(),
		cfg:    cfg,
	}
}

func (s *Server) Run() {

	s.engine.Run(":8080")
}
