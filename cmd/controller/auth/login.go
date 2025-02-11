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

type usersLogin struct {
	Email    string `json:"email" validate:"required,min=5,max=50,email"`
	Password string `json:"password" validate:"required,min=8,max=50"`
}

func ViewLogin(c *fiber.Ctx) error {
	return c.Render("auth/login", web.FiberTemp{
		MetaData:   web.GetMetadata(),
		PageTitle:  "Login",
		IsLoggedIn: false,
		Optional: map[string]string{
			`Csrf`:              c.Locals("csrfToken").(string),
			`RecapthcaSiteKey`:  os.Getenv("RECHAPTCHA_SITE_KEY"),
			"GoogleLoginStatus": os.Getenv("GOOGLE_LOGIN_STATUS"),
		},
	}, "layout/auth/main")
}

func ProcessLogin(c *fiber.Ctx) error {
	if os.Getenv("RECHAPTCHA_SECRET_KEY") != "" && !google_recaptcha.VerifyRecaptcha(c) {
		return c.JSON(helper.AjaxReturnError("reCAPTCHA verifikasi gagal, coba lagi"))
	}

	v, trans := helper.SetLanguageID()

	UsersLogin := usersLogin{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	Validation := v.Struct(UsersLogin)
	if Validation != nil {
		for _, err := range Validation.(validator.ValidationErrors) {
			return c.Status(500).JSON(helper.AjaxReturnError(err.Translate(*trans)))
		}
	}

	repo := repositoryimplementation.NewRepositoryUsers(configs.GetConnectionDB())
	err, getUser := repo.FindUsersByEmail(context.Background(), UsersLogin.Email)

	if err != nil {
		return c.Status(500).JSON(helper.AjaxReturnError(err.Error()))
	}

	//VERIFIKASI PASSWORD
	if !helper.VerifyPassword(UsersLogin.Password, getUser.PasswordHash) {
		return c.Status(500).JSON(helper.AjaxReturnError("Password anda salah"))
	}

	err = helper.SetUserSession(getUser, c)
	if err != nil {
		return c.Status(500).JSON(helper.AjaxReturnError(err.Error()))
	}

	return c.Status(200).JSON(helper.AjaxReturnSuccess("Login anda berhasil, terimakasih"))

}
