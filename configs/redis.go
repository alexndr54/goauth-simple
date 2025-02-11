package configs

import (
	"autentikasi1/cmd/helper/handle_panic"
	"github.com/gofiber/storage/redis/v3"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
)

func GetRedis() *redis.Storage {
	defer handle_panic.DeferFunc(logrus.FatalLevel, "Ada kesalahan redis/server belum aktif")

	var (
		redisHost      = os.Getenv("REDIS_HOST")
		redisUsername  = os.Getenv("REDIS_USERNAME")
		redisPassword  = os.Getenv("REDIS_PASSWORD")
		redisDatabase1 = os.Getenv("REDIS_DATABASE")
		redisPort1     = os.Getenv("REDIS_PORT")
	)

	if redisHost == "" || redisPort1 == "" {
		panic("REDIS_HOST & REDIS_PORT tidak boleh kosong,cek file .env")
	}

	redisPort, err := strconv.Atoi(redisPort1)
	handle_panic.PanicIfErr("REDIS_PORT hanya boleh angka,cek file .env", err)

	redisDatabase, err := strconv.Atoi(redisDatabase1)
	handle_panic.PanicIfErr("REDIS_DATABASE hanya boleh angka,default: 1,cek file .env", err)

	rd := redis.New(redis.Config{
		Host:      redisHost,
		Port:      redisPort,
		Username:  redisUsername,
		Password:  redisPassword,
		Database:  redisDatabase,
		Reset:     false,
		TLSConfig: nil,
		PoolSize:  10 * runtime.GOMAXPROCS(0),
	})

	return rd
}
