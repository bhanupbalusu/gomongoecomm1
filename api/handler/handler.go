package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bhanupbalusu/gomongoecomm1/domain/model"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Get(c *fiber.Ctx) error {
	productList, err := h.ProductService.Get()
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(productList)
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println(id)
	product, err := h.ProductService.GetByID(id)
	if err != nil {
		fmt.Println("------ Error from handler.ProductService.GetByID ----")
		log.Fatal(err)
	}
	return c.JSON(product)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	req := model.ProductModel{}
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	fmt.Println("------- Inside Handler Create Method Before Calling ProductService.Create -----------")
	fmt.Println(req)
	pid, err := h.ProductService.Create(req)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
			"data":    nil,
		})
		return err
	}
	fmt.Println("-------- Inside Handler Create Method After Calling ProductService.Create ----------")
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Product Created", "product_id": pid})
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	req := model.ProductModel{}
	fmt.Println("-----------api/handler.Update Before calling c.BodyParser ----------")
	if err := c.BodyParser(&req); err != nil {
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse body",
			"error":   err,
		})
	}

	fmt.Println("-----------api/handler.Update Before calling h.ProductService.Update ----------")

	if err := h.ProductService.Update(req, id); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Product failed to update",
			"error":   err.Error(),
		})
	}

	fmt.Println("-----------api/handler.Update Before calling final return----------")
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Product updated successfully",
	})
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.ProductService.Delete(id); err != nil {
		log.Fatal(err)
	}
	return c.SendString("Product Is Deleted") // send text
}
