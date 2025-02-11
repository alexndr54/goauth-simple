package auth

import (
	"autentikasi1/cmd/helper"
	"autentikasi1/cmd/model/web"
	repositoryimplementation "autentikasi1/cmd/repository/repository_impl"
	"autentikasi1/cmd/third-party/google-recaptcha"
	"autentikasi1/configs"
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"os"
)

type usersChangePassword struct {
	Password  string `json:"password" validate:"required,min=8,max=100"`
	Password2 string `json:"password2" validate:"required,eqfield=Password"`
	Token     string `json:"token" validate:"required,alphanum"`
}

func ViewChangePassword(c *fiber.Ctx) error {
	var repo = repositoryimplementation.NewRepositoryReset(configs.GetConnectionDB())
	Token := helper.GetAlphaNumeric(c.Params("token"))
	err, UserReset := repo.FindResetPasswordByToken(context.Background(), Token)
	if err != nil {
		return c.Render("auth/change-password", web.FiberTemp{
			MetaData:   web.GetMetadata(),
			PageTitle:  "Ubah Password",
			IsLoggedIn: false,
			Optional: map[string]string{
				`Csrf`:             c.Locals("csrfToken").(string),
				"Error":            err.Error(),
				`RecapthcaSiteKey`: os.Getenv("RECHAPTCHA_SITE_KEY"),
			},
		}, "layout/auth/main")
	}

	return c.Render("auth/change-password", web.FiberTemp{
		MetaData:   web.GetMetadata(),
		PageTitle:  "Ubah Password",
		IsLoggedIn: false,
		Optional: map[string]string{
			`Csrf`:  c.Locals("csrfToken").(string),
			"Email": UserReset.Email,
			"Token": UserReset.Token,
		},
	}, "layout/auth/main")
}

func ProcessChangePassword(c *fiber.Ctx) error {
	if os.Getenv("RECHAPTCHA_SECRET_KEY") != "" && !google_recaptcha.VerifyRecaptcha(c) {
		return c.JSON(helper.AjaxReturnError("reCAPTCHA verifikasi gagal, coba lagi"))
	}

	var repo = repositoryimplementation.NewRepositoryReset(configs.GetConnectionDB())
	v, trans := helper.SetLanguageID()

	UsersReset := usersChangePassword{
		Password:  c.FormValue("password1"),
		Password2: c.FormValue("password2"),
		Token:     c.Params("token"),
	}

	Validation := v.Struct(UsersReset)

	if Validation != nil {
		for _, err := range Validation.(validator.ValidationErrors) {
			return c.Status(500).JSON(helper.AjaxReturnError(err.Translate(*trans)))
		}
	}

	err, i := repo.FindResetPasswordByToken(context.Background(), UsersReset.Token)
	if err != nil {
		return c.Status(500).JSON(helper.AjaxReturnError(err.Error()))
	}

	err = repo.DeleteResetPasswordByEmail(context.Background(), i.Email)
	if err != nil {
		return c.Status(500).JSON(helper.AjaxReturnError(err.Error()))
	}

	PasswordHashed, err := helper.HashPassword(UsersReset.Password)
	if err != nil {
		return c.Status(500).JSON(helper.AjaxReturnError(err.Error()))
	}

	repos := repositoryimplementation.NewRepositoryUsers(configs.GetConnectionDB())
	err = repos.ChangePasswordByEmail(context.Background(), i.Email, PasswordHashed)
	if err != nil {
		return c.Status(500).JSON(helper.AjaxReturnError(err.Error()))
	}

	return c.Status(200).JSON(helper.AjaxReturnSuccess("Password anda berhasil di ubah, terimakasih"))

}
