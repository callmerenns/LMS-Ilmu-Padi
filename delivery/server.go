package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/kelompok-2/ilmu-padi/config"
)

// Note: All usecase / service should be injected
type Server struct {
	engine *gin.Engine
	cfg    *config.Config
}
