package test

import (
	"autentikasi1/cmd/helper"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGenerateApiUsername(t *testing.T) {
	str, err := helper.GenerateApiUsername()
	assert.Nil(t, err)
	t.Log("API Username: ", str)
}

func TestGenerateApiKey(t *testing.T) {
	err, str := helper.GenerateApiKey()
	assert.Nil(t, err)
	t.Log("API Key: ", str)
}

func TestWaktu(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	waktu1 := time.Date(2025, 02, 9, 18, 05, 48, 0, loc).Add(10 * time.Minute)
	fmt.Println("WAKTU 1 BELUM LEWAT JAM SAAT INI")
	fmt.Println("waktu1.After: ", time.Now().After(waktu1))   //true
	fmt.Println("waktu1.Before: ", time.Now().Before(waktu1)) //false
	fmt.Println()

	waktu2 := time.Date(2025, 02, 10, 11, 00, 48, 0, loc).Add(10 * time.Minute)
	fmt.Println("waktu1.After: ", time.Now().After(waktu2))   //false
	fmt.Println("waktu1.Before: ", time.Now().Before(waktu2)) //true
}
