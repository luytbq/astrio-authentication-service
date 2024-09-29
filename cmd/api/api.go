package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/luytbq/astrio-authentication-service/internal/auth"
)

type Server struct {
	Port   string
	Prefix string
	DB     *sql.DB
}

func NewServer(port string, prefix string, db *sql.DB) *Server {
	return &Server{
		Port:   port,
		Prefix: prefix,
		DB:     db,
	}
}

func (s *Server) Run() error {
	engine := gin.Default()

	engine.Use(CORSMiddleware())
	repo := auth.NewRepo(s.DB)
	handler := auth.NewHandler(repo)

	handler.RegisterRoutes(engine)

	err := engine.Run(":" + s.Port)

	return err
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
