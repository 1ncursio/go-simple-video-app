package main

import (
	"context"
	"fmt"
	"log"

	"github.com/1ncursio/go-simple-video-app/ent"
	"github.com/1ncursio/go-simple-video-app/ent/user"
	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.Name("a8m")).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("filed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
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
