package configs

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/fiber/v2/utils"
	"time"
)

func GetSession() *session.Store {
	sess := session.New(session.Config{
		Expiration:   24 * time.Hour,
		Storage:      GetRedis(),
		KeyLookup:    "cookie:session_id",
		KeyGenerator: utils.UUIDv4,
	})

	return sess
}
