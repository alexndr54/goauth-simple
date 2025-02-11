package google_recaptcha

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"net/url"
	"os"
)

func VerifyRecaptcha(c *fiber.Ctx) bool {
	recaptchaResponse := c.FormValue("g-recaptcha-response")
	verifyURL := "https://www.google.com/recaptcha/api/siteverify"

	params := url.Values{
		"secret":   {os.Getenv("RECHAPTCHA_SECRET_KEY")},
		"response": {recaptchaResponse},
		"remoteip": {c.IP()},
	}
	resp, err := http.PostForm(verifyURL, params)
	if err != nil {
		return false
	}

	defer resp.Body.Close()

	// Parse the JSON response
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return false
	}

	// Check if the reCAPTCHA verification was successful
	fmt.Println("Data:", data)
	if data["success"].(bool) && data["score"].(float64) >= 0.6 {
		return true
	} else {
		return false
	}
}
