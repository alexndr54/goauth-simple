package configs

import (
	"autentikasi1/cmd/helper/handle_panic"
	"autentikasi1/configs/configs_logrus"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"time"
)

func GetConnectionDB() *sql.DB {
	defer handle_panic.DeferFunc(logrus.FatalLevel, "Kesalahan konfigurasi database di file .env, atau ada yang kosong")

	var (
		dbHost                  = os.Getenv("DB_HOST")
		dbPort                  = os.Getenv("DB_PORT")
		dbUser                  = os.Getenv("DB_USERNAME")
		dbPass                  = os.Getenv("DB_PASSWORD")
		dbName                  = os.Getenv("DB_NAME")
		valdbSetMaxIdleConns    = os.Getenv("DB_SetMaxIdleConns")
		valdbSetMaxOpenConns    = os.Getenv("DB_SetMaxOpenConns")
		valdbSetConnMaxIdleTime = os.Getenv("DB_SetConnMaxIdleTime")
		valdbSetConnMaxLifetime = os.Getenv("DB_SetConnMaxLifetime")
		Logging                 = configs_logrus.SetupLogger()
	)

	if dbHost == "" || dbPort == "" || dbUser == "" || dbName == "" || valdbSetMaxIdleConns == "" || valdbSetMaxOpenConns == "" || valdbSetConnMaxIdleTime == "" || valdbSetConnMaxLifetime == "" {
		Logging.WithFields(logrus.Fields{
			"host":     dbHost,
			"port":     dbPort,
			"user":     dbUser,
			"name":     dbName,
			"maxIdle":  valdbSetMaxIdleConns,
			"maxOpen":  valdbSetMaxOpenConns,
			"idle":     valdbSetConnMaxIdleTime,
			"lifetime": valdbSetConnMaxLifetime,
		}).Info("Blank Value")
		panic("Periksa lagi pengaturan database di file .env, Pastikan tidak ada yang kosong, atau tidak sesuai yang di minta")
	}

	dbSetMaxIdleConns, err := strconv.Atoi(valdbSetMaxIdleConns)
	handle_panic.PanicIfErr("DB_SetMaxIdleConns hanya boleh angka", err)

	dbSetMaxOpenConns, err := strconv.Atoi(valdbSetMaxOpenConns)
	handle_panic.PanicIfErr("DB_SetMaxOpenConns hanya boleh angka", err)

	dbSetConnMaxIdleTime1, err := strconv.Atoi(valdbSetConnMaxIdleTime)
	dbSetConnMaxIdleTime := time.Duration(dbSetConnMaxIdleTime1) * time.Minute
	handle_panic.PanicIfErr("DB_SetConnMaxIdleTime hanya boleh angka", err)

	dbSetConnMaxLifetime1, err := strconv.Atoi(valdbSetConnMaxLifetime)
	dbSetConnMaxLifetime := time.Duration(dbSetConnMaxLifetime1) * time.Hour
	handle_panic.PanicIfErr("DB_SetConnMaxLifetime hanya boleh angka", err)

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(dbSetMaxIdleConns)
	db.SetMaxOpenConns(dbSetMaxOpenConns)
	db.SetConnMaxIdleTime(dbSetConnMaxIdleTime)
	db.SetConnMaxLifetime(dbSetConnMaxLifetime)

	return db

}
