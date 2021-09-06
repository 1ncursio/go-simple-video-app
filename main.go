package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"statusCode": fiber.StatusCreated,
			"message":    "TEST",
			"data":       fiber.Map{},
		})
	})

	v1 := app.Group("/api/v1", func(c *fiber.Ctx) error {
		return c.Next()
	})

	videoRouter := v1.Group("/videos", func(c *fiber.Ctx) error {
		return c.Next()
	})

	// GET /api/v1/videos
	videoRouter.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{})
	})

	// GET /api/v1/videos/1
	videoRouter.Get("/:videoId", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{})
	})

	// POST /api/v1/videos
	videoRouter.Post("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{})
	})

	// PATCH /api/v1/videos/1
	videoRouter.Patch("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{})
	})

	// DELETE /api/v1/videos/1
	videoRouter.Delete("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{})
	})

	log.Fatal(app.Listen("localhost:3000"))
}
