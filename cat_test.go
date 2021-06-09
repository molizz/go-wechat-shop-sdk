package shop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCat_Get(t *testing.T) {
	at, err := NewAccessToken(APPID, SECRETE).Get()
	assert.Nil(t, err)

	result, err := NewCat(at).Get()
	assert.Nil(t, err)
	assert.NotEmpty(t, result.ThirdCatList)
}
