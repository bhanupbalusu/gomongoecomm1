package api

import (
	s "github.com/bhanupbalusu/gomongoecomm1/domain/interface/service"

	"github.com/gofiber/fiber/v2"
)

type RedirectHandler interface {
	Get(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type Handler struct {
	ProductService s.ProductServiceInterface
}

func NewHandler(productService s.ProductServiceInterface) RedirectHandler {
	return &Handler{ProductService: productService}
}
