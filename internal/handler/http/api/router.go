package api

import "github.com/gofiber/fiber/v3"

type Handler interface {
	AddRoutes(r fiber.Router)
}

type Router struct {
	router *fiber.App
}

func NewRouter() *Router {
	return &Router{router: fiber.New()}
}

func (r *Router) WithHandler(h Handler) *Router {
	h.AddRoutes(r.router)

	return r
}
