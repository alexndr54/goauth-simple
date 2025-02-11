package middleware

import (
	"autentikasi1/cmd/helper"
	"github.com/gofiber/fiber/v2"
)

func HasBenLogin(c *fiber.Ctx) error {
	err, user := helper.GetUserSession(c)
	if err != nil {
		return c.Redirect("/auth/login")
	}

	c.Locals("users", user)
	return c.Next()
}
