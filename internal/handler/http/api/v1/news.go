package v1

import (
	"encoding/json"
	"github.com/SemyonTolkachyov/news-api/internal/model/input"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

// getNews get news paged
func (h *Handler) getNews(c fiber.Ctx) error {
	size, err := strconv.Atoi(c.Query("size"))
	if err != nil {
		size = 20
	}
	number, err := strconv.Atoi(c.Query("number"))
	if err != nil {
		number = 1
	}
	all, err := h.services.NewsService.GetPaged(c.Context(), size, number)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return err
	}
	err = c.JSON(&all)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return err
	}
	return nil
}

// editNews update news with categories
func (h *Handler) editNews(c fiber.Ctx) error {
	var inp input.UpdateNews
	err := json.Unmarshal(c.Body(), &inp)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}
	err = h.validate.Struct(inp)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}
	newsId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	err = h.services.NewsService.Update(c.Context(), newsId, inp)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return err
	}
	c.Status(fiber.StatusOK)
	return nil
}
