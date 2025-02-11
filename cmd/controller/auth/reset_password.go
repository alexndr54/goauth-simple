package auth

import (
	"autentikasi1/cmd/email"
	"autentikasi1/cmd/helper"
	"autentikasi1/cmd/model/internal_db"
	"autentikasi1/cmd/model/web"
	repositoryimplementation "autentikasi1/cmd/repository/repository_impl"
	"autentikasi1/cmd/third-party/google-recaptcha"
	"autentikasi1/configs"
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"os"
)

type usersReset struct {
	Email string `json:"email" validate:"required,email,min=5,max=30"`
}

func ViewResetPassword(c *fiber.Ctx) error {
	return c.Render("auth/reset", web.FiberTemp{
		MetaData:   web.GetMetadata(),
		PageTitle:  "Reset Password",
		IsLoggedIn: false,
		Optional: map[string]string{
			`Csrf`:             c.Locals("csrfToken").(string),
			`RecapthcaSiteKey`: os.Getenv("RECHAPTCHA_SITE_KEY"),
		},
	}, "layout/auth/main")
}

func ProcessResetPassword(c *fiber.Ctx) error {
	if os.Getenv("RECHAPTCHA_SECRET_KEY") != "" && !google_recaptcha.VerifyRecaptcha(c) {
		return c.JSON(helper.AjaxReturnError("reCAPTCHA verifikasi gagal, coba lagi"))
	}
	v, trans := helper.SetLanguageID()

	users := usersReset{Email: c.FormValue("email")}

	Validation := v.Struct(users)

	if Validation != nil {
		for _, err := range Validation.(validator.ValidationErrors) {
			return c.Status(500).JSON(helper.AjaxReturnError(err.Translate(*trans)))
		}
	}

	tokenReset, errToken := helper.GenerateRandomString(200)
	if errToken != nil {
		return c.Status(500).JSON(helper.AjaxReturnError("Gagal membuat token"))
	}

	repo := repositoryimplementation.NewRepositoryReset(configs.GetConnectionDB())
	err := repo.CreateResetPassword(context.Background(), &internal_db.ModelResetPassword{
		Email: users.Email,
		Token: tokenReset,
	})

	if err != nil {
		return c.Status(500).JSON(helper.AjaxReturnError(err.Error()))
	}

	err = email.SendMailReset(users.Email, tokenReset)
	if err != nil {
		return c.Status(500).JSON(helper.AjaxReturnError(err.Error()))
	}

	return c.Status(200).JSON(helper.AjaxReturnSuccess("Email telah dikirim,Cek spam/inbox"))
}
