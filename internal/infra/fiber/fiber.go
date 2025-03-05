package fiber

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const idleTimeout = 5 * time.Second

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})

	app.Use(cors.New())

	return app
}
