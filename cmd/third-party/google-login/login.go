package google_login

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type GoogleUserInfo struct {
	ID            int    `json:"id"`
	Email         string `json:"email"`
	Fullname      string `json:"fullname"`
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
}

var (
	ClientID       = os.Getenv("GOOGLE_CLIENT_ID")
	SecretID       = os.Getenv("GOOGLE_CLIENT_SECRET")
	CallbackGoogle = os.Getenv("SITE_URL") + "auth/login/google/callback"
)

func init() {
	goth.UseProviders(
		google.New(ClientID, SecretID, CallbackGoogle, "email", "profile"),
	)
}

const (
	GoogleAuthURL     = "https://accounts.google.com/o/oauth2/auth"
	GoogleTokenURL    = "https://oauth2.googleapis.com/token"
	GoogleUserInfoURL = "https://www.googleapis.com/oauth2/v2/userinfo"
)

func GetGoogleLoginURL() string {
	params := url.Values{}
	params.Add("client_id", ClientID)
	params.Add("redirect_uri", CallbackGoogle)
	params.Add("response_type", "code")
	params.Add("scope", "email profile")
	params.Add("state", "state-token")

	authURL := GoogleAuthURL + "?" + params.Encode()
	return authURL
}

func GetGoogleUserInfo(code string) (error, *GoogleUserInfo) {
	if code == "" {
		return errors.New("ada kesalahan, kode autentikasi tidak ditemukan"), nil
	}

	data := url.Values{}
	data.Set("client_id", ClientID)
	data.Set("client_secret", SecretID)
	data.Set("redirect_uri", CallbackGoogle)
	data.Set("grant_type", "authorization_code")
	data.Set("code", code)

	req, err := http.NewRequest("POST", GoogleTokenURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return errors.New("ada kesalahan, tidak bisa membuat request"), nil
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("ada kesalahan, tidak bisa mengirim request"), nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New("ada kesalahan, tidak bisa membaca response"), nil
	}

	var tokenResp map[string]interface{}
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return errors.New("ada kesalahan,  tidak bisa decode json"), nil
	}

	accessToken, ok := tokenResp["access_token"].(string)
	if !ok {
		//tokenResp["error_description"].(string)
		return errors.New("Silahkan melakukan login ulang"), nil
	}

	req, err = http.NewRequest("GET", GoogleUserInfoURL, nil)
	if err != nil {
		return errors.New("Ada kesalahan, tidak bisa membuat request"), nil
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err = client.Do(req)
	if err != nil {
		return errors.New("Ada kesalahan, tidak bisa mengirim request"), nil
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Ada kesalahan, tidak bisa membaca response"), nil
	}

	var userInfo map[string]interface{}
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return errors.New("Ada kesalahan, tidak bisa decode json"), nil
	}

	ID, _ := strconv.Atoi(userInfo["id"].(string))
	user := GoogleUserInfo{
		ID:            ID,
		Email:         userInfo["email"].(string),
		Fullname:      userInfo["given_name"].(string),
		Picture:       userInfo["picture"].(string),
		VerifiedEmail: userInfo["verified_email"].(bool),
	}
	return nil, &user
}
