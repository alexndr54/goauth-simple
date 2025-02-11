# Goauth-simple
### Merupakan aplikasi berbasis website yang dimana mengunakan beberapa teknologi:
- Go
- Gofiber (Framework)
- Mysql (Database)
- Redis   

## Fitur
- Login
- Register
- Reset Password (Dengan email smtp)
- Change Password
- Login Via Google

## Persyaratan
1. Minimal golang versi `1.23.4`
2. Redis

## Install
1. Silahkan buat folder dahulu,kemudian buka folder tersebut di code editor, kemudian buka juga di terminal.
2. kemudian kita melakukan clone,   
`git clone https://github.com/alexndr54/goauth-simple`
3. Silahkan anda ubah nama file `.env.example` menjadi `.env` dan atur yang diperlukan disana   
4. Kemudian WAJIB migrasi table terlebih dahulu dengan cara   
`migrate -path=migrations/db -database="mysql://username:password@(localhost:3306)/nama_database" up`   
Ganti:   
`username`: dengan username database,misal root   
`password`: password database   
`nama_database`: dengan nama databasenya,misal autentikasi

## Cara Menggunakan
1. Buka 2 terminal
2. terminal 1: `redis` untuk menjalankan redis
3. terminal 2: `go run main.go` untuk menjalankan gofiber/go

### Cara Mengubah Metadata website
Apa itu metadata? metadata disini berisi seperti judul website, meta tag html,favicon path,logo path,yang digunakan akan ditampilkan dihalaman web
1. Buka dan sesuaikan file metadata.json

## Algoritma Reset Password
1. User di haruskan register dahulu
2. User melakukan request link untuk merubah password dihalaman reset password,Dengan masa aktif link hanya 10 menit
3. user mendapatkan link untuk mengubah password, Jika expired user di minta request link ulang
4. User melakukan ubah password.
5. User berhasil ubah password, bisa langsung login

## Struktur Folder
1. `cmd`    
    - `controller` berisi file yang menagani view halaman website yang menangani proses
    - `email` berisi fungsi untuk mengirim email,di sediakan juga function beserta template email
    - `helper` berisi fungsi fungsi sebagai pembantu, seperti Hash Password, Verifikasi password
    - `middleware` berisi middleware
    - `model` berisi kumpulan struct dari table
    - `repository` kumpulan kode untuk menangani proses golang - database
    - `third-party` proses pihak ketiga disini, Seperti google recaptcha, request payment gateway
2. `configs` berisi pengaturan seperti membuat function untuk menghubungkan golang ke database mysql, ke redis, ke session
3. `migrations/db` berisi file `.sql` yang digunakan untuk migrate
4. `init` folder yang berisi fungsi init
5. `test` kumpulan testing
6. `web` kumpulan file .html

## ROUTE
1. `yourdomain.com/auth/login` login
2. `yourdomain.com/auth/register` register account
3. `yourdomain.com/auth/reset-password` reset password
4. `yourdomain.com/auth/change-password/{your token}` ubah password
5. `yourdomain.com/auth/login/google/` login via google