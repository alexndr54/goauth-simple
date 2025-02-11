package test

import (
	"autentikasi1/cmd/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnkripsi(t *testing.T) {
	res, err := helper.MyEncrypt([]byte("arman123"))
	assert.Nil(t, err)
	t.Log("Enkripsi", res)

	decrypt, err := helper.MyDecrypt(res)
	assert.Nil(t, err)
	t.Log("Dekripsi", decrypt)
}
