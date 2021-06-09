package shop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister_Check(t *testing.T) {
	at, err := NewAccessToken(APPID, SECRETE).Get()
	assert.Nil(t, err)

	r := NewRegister(at)

	result, err := r.Check()
	assert.Nil(t, err)
	assert.NotEmpty(t, result.Data)
	assert.NotEmpty(t, result.Data.AccessInfo.SpuAuditSuccess)
}

func TestRegister_Apply(t *testing.T) {
	at, err := NewAccessToken(APPID, SECRETE).Get()
	assert.Nil(t, err)

	r := NewRegister(at)

	_, err = r.Apply()
	assert.Nil(t, err)
}
