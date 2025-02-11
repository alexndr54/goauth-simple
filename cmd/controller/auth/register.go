package auth

import (
	"autentikasi1/cmd/helper"
	"autentikasi1/cmd/model/internal_db"
	"autentikasi1/cmd/model/web"
	repositoryimplementation "autentikasi1/cmd/repository/repository_impl"
	"autentikasi1/cmd/third-party/google-recaptcha"
	"autentikasi1/configs"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"os"
)

type usersRegister struct {
	Fullname    string `json:"fullname" validate:"required,min=5,max=20"`
	Email       string `json:"email" validate:"required,email,min=5,max=30"`
	Password    string `json:"password" validate:"required,min=8,max=100"`
	Password2   string `json:"password2" validate:"required,eqfield=Password"`
	ApiUsername string `json:"api_username" validate:"required"`
	ApiKey      string `json:"api_key" validate:"required"`
}

func ViewRegister(c *fiber.Ctx) error {
	return c.Render("auth/register", web.FiberTemp{
		MetaData:   web.GetMetadata(),
		PageTitle:  "Register",
		IsLoggedIn: false,
		Optional: map[string]string{
			`Csrf`:             c.Locals("csrfToken").(string),
			`RecapthcaSiteKey`: os.Getenv("RECHAPTCHA_SITE_KEY"),
		},
	}, "layout/auth/main")
}

func ProcessRegister(c *fiber.Ctx) error {
	if os.Getenv("RECHAPTCHA_SECRET_KEY") != "" && !google_recaptcha.VerifyRecaptcha(c) {
		return c.JSON(helper.AjaxReturnError("reCAPTCHA verifikasi gagal, coba lagi"))
	}

	v, trans := helper.SetLanguageID()

	err, Apikey := helper.GenerateApiKey()
	if err != nil {
		return c.Status(500).JSON(helper.AjaxReturnError(err.Error()))
	}
	ApiUsername, err := helper.GenerateApiUsername()
	if err != nil {
		return c.Status(500).JSON(helper.AjaxReturnError(err.Error()))
	}

	users := usersRegister{
		Fullname:    c.FormValue("fullname"),
		Email:       c.FormValue("email"),
		Password:    c.FormValue("password1"),
		Password2:   c.FormValue("password2"),
		ApiUsername: ApiUsername,
		ApiKey:      Apikey,
	}

	Validation := v.Struct(users)

	if Validation != nil {
		for _, err := range Validation.(validator.ValidationErrors) {
			return c.Status(500).JSON(helper.AjaxReturnError(err.Translate(*trans)))
		}
	}

	passwordHashed, err := helper.HashPassword(users.Password)
	if err != nil {
		return c.Status(500).JSON(helper.AjaxReturnError(err.Error()))
	}
	repo := repositoryimplementation.NewRepositoryUsers(configs.GetConnectionDB())
	err = repo.CreateUsers(context.Background(), &internal_db.ModelUsers{
		Fullname: users.Fullname,
		Email:    users.Email,
		ApiUsername: sql.NullString{
			String: users.ApiUsername,
			Valid:  true,
		},
		ApiKey: sql.NullString{
			String: users.ApiKey,
			Valid:  true,
		},
		PasswordHash: passwordHashed,
	})

	if err != nil {
		return c.Status(500).JSON(helper.AjaxReturnError(err.Error()))
	}

	return c.Status(200).JSON(helper.AjaxReturnSuccess("Pendaftaran berhasil silahkan login"))
}
