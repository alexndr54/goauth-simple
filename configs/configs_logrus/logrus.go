package configs_logrus

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func SetupLogger() *logrus.Logger {
	log := logrus.New()

	// Setup log rotation
	log.SetOutput(&lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    10,   // Maksimal ukuran file log dalam MB
		MaxBackups: 5,    // Maksimal file log yang disimpan sebelum dihapus
		MaxAge:     30,   // Berapa hari file log disimpan sebelum dihapus
		Compress:   true, // Mengompresi log lama
	})

	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)

	return log
}
