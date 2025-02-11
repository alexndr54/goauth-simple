package configs

import (
	"github.com/gofiber/template/html/v2"
)

func GetEngineTemplate() *html.Engine {
	return html.New("./web", ".html")
}
