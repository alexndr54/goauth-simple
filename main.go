package main

import (
	auth2 "autentikasi1/cmd/controller/auth"
	app "autentikasi1/cmd/controller/dashboard"
	"autentikasi1/cmd/helper/handle_panic"
	"autentikasi1/cmd/middleware"
	"autentikasi1/configs"
	_ "autentikasi1/init"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	gofiber := configs.GetGoFiber()
	gofiber.Use(csrf.New(csrf.Config{
		ContextKey: "csrfToken",
		//KeyLookup:  "header: X-Csrf-Token",
		//CookieName: "__Host-csrf_",
		//CookieSecure:      true,
		//CookieSessionOnly: true,
		//CookieHTTPOnly:    true,
		//Expiration:   1 * time.Hour,
		//KeyGenerator: utils.UUIDv4,
		Session: configs.GetSession(),
		//SessionKey:        "fiber.csrf.token",
		//HandlerContextKey: "fiber.csrf.handler",
		SingleUseToken: false, // <--- set single use token to false
	}))
	gofiber.Use(logger.New(logger.Config{
		Format:     "[${time}]:[${ip}]:[${port}] ${status} - ${method} ${path}\n",
		TimeFormat: "02-01-2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	gofiber.Static("/public", "./public")

	//DASHBOARD
	dashboard := gofiber.Group("/app/")
	dashboard.Use(middleware.HasBenLogin)
	dashboard.Get("home", app.ViewHome)

	//AUTENTIKASI
	auth := gofiber.Group("/auth/")
	auth.Get("login", auth2.ViewLogin)
	auth.Get("login/google/", auth2.GetGoogleLoginURL)
	auth.Get("login/google/callback", auth2.ProcessGoogleLogin)

	auth.Get("register", auth2.ViewRegister)
	auth.Get("reset-password", auth2.ViewResetPassword)
	auth.Get("change-password/:token", auth2.ViewChangePassword)

	//AJAX
	ajax := gofiber.Group("/ajax/")
	ajax.Post("auth/login", auth2.ProcessLogin)
	ajax.Post("auth/register", auth2.ProcessRegister)
	ajax.Post("auth/reset-password", auth2.ProcessResetPassword)
	ajax.Post("auth/change-password/:token", auth2.ProcessChangePassword)

	//LISTEN
	err := gofiber.Listen(":1000")
	handle_panic.PanicIfErr("Terjadi kesalahan ketika listen gofiber", err)
}
