package helper

import "time"

// Mengembalikan time.zone jakarta dan waktu menggunakan timezone jakarta
func GetDateTime(times time.Time) (*time.Location, time.Time) {
	jakarta, _ := time.LoadLocation("Asia/Jakarta")
	return jakarta, time.Date(times.Year(), times.Month(), times.Day(), times.Hour(), times.Minute(), times.Second(), 0, jakarta)
}
