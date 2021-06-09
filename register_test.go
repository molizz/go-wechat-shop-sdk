package shop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister_Check(t *testing.T) {
	at, err := NewAccessToken(APPID, SECRETE).Get()
	assert.Nil(t, err)

	r := NewRegister(at)

	_, err = r.Check()
	assert.Nil(t, err)
}
