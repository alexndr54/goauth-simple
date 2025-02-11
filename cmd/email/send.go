package email

import (
	"fmt"
	"os"
	"time"
)

func SendMailReset(emailTujuan, tokenReset string) error {
	link := os.Getenv("SITE_URL") + fmt.Sprintf("auth/change-password/%s", tokenReset)
	data := map[string]interface{}{
		"Link": link,
		"Waktu": time.Now().In(func() *time.Location {
			loc, _ := time.LoadLocation("Asia/Jakarta")
			return loc
		}(),
		).Format("02 Jan 2006 15:04:00"),
	}
	html := `
	<html>
	<head>
		<title>Reset Password</title>
	</head>
	<body>
		<h2><b>Reset Password</b></h2>
		<p>Kami menerima permintaan untuk reset password pada akun anda, Jika anda merasa melakukan permintaan reset password, silahkan anda klik tombol dibawah ini untuk mengubahnya.</p>
		<a href="{{ .Link }}">Klik Disini</a>
		<p>Jika anda merasa tidak melakukan permintaan perubahan password, silahkan abaikan saja</p>
		<p>Terimakasih</p><br>
		<p>Waktu: {{ .Waktu }}</p>
	</body>
	</html>
	`

	err := SendHTMLEmail(emailTujuan, html, "Permintaan Reset Password", data)
	if err != nil {
		return err
	}

	return nil
}
func SendPassword(emailTujuan, Password string) error {
	data := map[string]interface{}{
		"Waktu": time.Now().In(func() *time.Location {
			loc, _ := time.LoadLocation("Asia/Jakarta")
			return loc
		}(),
		).Format("02 Jan 2006 15:04:00"),
		"Password": Password,
	}
	html := `
	<html>
	<head>
		<title>Pendaftaran Berhasil</title>
	</head>
	<body>
		<h2><b>Password Sementara</b></h2>
		<p>Hai, kami sudah menerima permintaan pendaftaran anda, anda sekarang bisa login menggunakan akun google maupun email dan password, Nah berikutn kami lampirkan password sementara anda.</p>
		<p>Silahkan segera mengubah password anda dalam 24 jam  ya</p>
		<p><h2>Password: {{ .Password }}</h2></p>
		<p>Jika anda merasa tidak melakukan permintaan perubahan password, silahkan abaikan saja</p>
		<p>Terimakasih</p><br>
		<p>Waktu: {{ .Waktu }}</p>
	</body>
	</html>
	`

	err := SendHTMLEmail(emailTujuan, html, "Permintaan Reset Password", data)
	if err != nil {
		return err
	}

	return nil
}
