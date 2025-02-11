package auth

import (
	"autentikasi1/cmd/email"
	"autentikasi1/cmd/helper"
	"autentikasi1/cmd/model/internal_db"
	"autentikasi1/cmd/model/web"
	repositoryimplementation "autentikasi1/cmd/repository/repository_impl"
	google_login "autentikasi1/cmd/third-party/google-login"
	"autentikasi1/configs"
	"context"
	"github.com/gofiber/fiber/v2"
	"os"
)

var optional map[string]string

func GetGoogleLoginURL(c *fiber.Ctx) error {
	GOOGLE_LOGIN_STATUS := os.Getenv("GOOGLE_LOGIN_STATUS")
	if GOOGLE_LOGIN_STATUS == "false" {
		optional = map[string]string{
			"error": "Mohon maaf kami sedang menonaktifkan login via google",
		}

		return c.Render("auth/sosial-login", web.FiberTemp{
			MetaData:   web.GetMetadata(),
			PageTitle:  "Login Via Google",
			IsLoggedIn: false,
			Optional:   nil,
		}, "layout/auth/main")
	} else {
		return c.Redirect(google_login.GetGoogleLoginURL())
	}

}

func ProcessGoogleLogin(c *fiber.Ctx) error {

	GOOGLE_LOGIN_STATUS := os.Getenv("GOOGLE_LOGIN_STATUS")
	if GOOGLE_LOGIN_STATUS == "true" {
		err, g := google_login.GetGoogleUserInfo(c.Query("code"))
		if err == nil {
			Password, err := helper.GenerateRandomString(30)
			if err == nil {
				repo := repositoryimplementation.NewRepositoryUsers(configs.GetConnectionDB())
				err, getUser := repo.FindUsersOrCreateUsers(context.Background(), Password, &internal_db.ModelUsers{
					Fullname:      g.Fullname,
					Email:         g.Email,
					LoginProvider: "GOOGLE",
				})

				if err != nil {
					optional = map[string]string{
						"error": err.Error(),
					}
				} else if getUser != nil && getUser.LoginProvider == "REGISTER" {
					optional = map[string]string{
						"error": "Silahkan login menggunakan email dan password",
					}
				} else if getUser != nil && getUser.LoginProvider != "GOOGLE" {
					optional = map[string]string{
						"error": "Anda login menggunakan " + getUser.LoginProvider,
					}
				} else {

					err := helper.SetUserSession(getUser, c)
					if err != nil {
						optional = map[string]string{
							"error": err.Error(),
						}
					} else {
						optional = map[string]string{
							"email": getUser.Email,
							"name":  getUser.Fullname,
						}

						if getUser.PasswordHash == "true" {
							_ = email.SendPassword(getUser.Email, Password)
						}
					}

				}
			} else {
				optional = map[string]string{
					"error": err.Error(),
				}
			}
		} else {
			optional = map[string]string{
				"error": err.Error(),
			}
		}
	} else {
		optional = map[string]string{
			"error": "Mohon maaf kami sedang menonaktifkan login via google",
		}
	}

	return c.Render("auth/sosial-login", web.FiberTemp{
		MetaData:   web.GetMetadata(),
		PageTitle:  "Login Via Google",
		IsLoggedIn: false,
		Optional:   optional,
	}, "layout/auth/main")
}
