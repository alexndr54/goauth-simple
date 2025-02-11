package test

import (
	"autentikasi1/cmd/helper"
	"testing"
)

func TestGetAlphaNumeric(t *testing.T) {
	str := `'"!@#$%^&*()_+{}|:?><tololg`
	res := helper.GetAlphaNumeric(str)
	t.Log(res)
}
