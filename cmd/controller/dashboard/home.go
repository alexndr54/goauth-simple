package dashboard

import (
	"autentikasi1/cmd/model/internal_db"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func ViewHome(c *fiber.Ctx) error {
	users := c.Locals("users").(*internal_db.ModelUsers)
	return c.SendString(fmt.Sprintf("Halo %s,Email kamu: %s", users.Fullname, users.Email))
}
