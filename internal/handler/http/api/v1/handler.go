package v1

import (
	"github.com/SemyonTolkachyov/news-api/internal/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	services *service.Service
	validate *validator.Validate
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services, validate: validator.New()}
}

func (h *Handler) AddRoutes(r fiber.Router) {
	h.setNewsRouters(r)
}

func (h *Handler) setNewsRouters(r fiber.Router) {
	r.Get("/list", h.getNews)
	r.Post("/edit/:Id", h.editNews)
}
