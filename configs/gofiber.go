package configs

import (
	"github.com/gofiber/fiber/v2"
)

func GetGoFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		Views:         GetEngineTemplate(),
	})

	return app
}
