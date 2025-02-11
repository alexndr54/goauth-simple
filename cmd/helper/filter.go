package helper

import "regexp"

func GetAlphaNumeric(str string) string {
	return regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(str, "")
}
