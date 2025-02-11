package helper

import (
	"autentikasi1/cmd/model/internal_db"
	"autentikasi1/configs"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func SetUserSession(user *internal_db.ModelUsers, c *fiber.Ctx) error {
	store, err := configs.GetSession().Get(c)
	if err != nil {
		return errors.New("Terjadi kesalahan pada pengaturan session")
	}

	JsonByte, err := json.Marshal(*user)
	if err != nil {
		return errors.New("Gagal mengubah ke bentuk json")
	}

	store.Set("users", JsonByte)
	err = store.Save()
	if err != nil {
		return errors.New("Gagal menyimpan session")
	}

	return nil
}

func GetUserSession(c *fiber.Ctx) (error, *internal_db.ModelUsers) {
	store, err := configs.GetSession().Get(c)
	if err != nil {
		return errors.New("Ada kesalahan pada pengaturan session"), nil
	}

	users := store.Get("users")
	if users == nil {
		return errors.New("Pengguna tidak ditemukan"), nil
	}

	var user internal_db.ModelUsers
	err = json.Unmarshal(users.([]byte), &user)
	if err != nil {
		return errors.New("Gagal mengubah ke bentuk json"), nil
	}

	return nil, &user
}
