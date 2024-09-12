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

	repo := auth.NewRepo(s.DB)
	handler := auth.NewHandler(repo)

	handler.RegisterRoutes(engine)

	err := engine.Run(":" + s.Port)

	return err
}
