package api

import (
	"errors"
	"fmt"
	"github.com/SemyonTolkachyov/news-api/internal/config"
	"github.com/gofiber/fiber/v3"
	"net/http"
)

type Server struct {
	srv *fiber.App
}

func NewServer(cfg config.Config) *Server {
	server := &Server{fiber.New(fiber.Config{
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		AppName:      cfg.Name,
	})}

	return server
}

func (s *Server) RegisterRoutes(r *Router) {
	s.srv = r.router
}

func (s *Server) Start(cfg config.Config) error {
	err := s.srv.Listen(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	return s.srv.Shutdown()
}
