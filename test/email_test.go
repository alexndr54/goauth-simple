package test

import (
	"autentikasi1/cmd/email"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendMail(t *testing.T) {
	err := godotenv.Load("../.env")
	assert.Nil(t, err)

	data := map[string]interface{}{
		"Nama":  "John Doe",
		"Email": "sialexsofficial@gmail.com",
		"Usia":  30,
	}

	err = email.SendHTMLEmail(data["Email"].(string), "Testing", `<!DOCTYPE html>
<html>
<head>
	<title>Email Test</title>
</head>
<body>
	<h1>Hai, {{.Nama}}</h1>
	<p>Terima kasih telah bergabung dengan kami. Berikut adalah informasi tentang diri Anda:</p>
	<p>Nama: {{.Nama}}</p>
	<p>Email: {{.Email}}</p>
	<p>Usia: {{.Usia}}</p>
</body>
</html>`, data)
	assert.Nil(t, err)
}
