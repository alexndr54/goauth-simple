package success

import (
	"autentikasi1/configs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	DB := configs.GetConnectionDB()
	assert.Nil(t, DB.Ping())
	t.Log("Database terkoneksi dengan baik")
}

func TestGetSession(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fatalf("Ada kesalahan pada redis %s", err)
		}
	}()

	session := configs.GetSession()
	assert.NotNil(t, session)

	t.Log("Koneksi session dengan redis berjalan dengan baik")
}
