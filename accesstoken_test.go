package shop

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	APPID   = os.Getenv("WEIXIN_MINIPROGRESS_APPID")
	SECRETE = os.Getenv("WEIXIN_MINIPROGRESS_SECRET")
)

func TestAccessToken_Get(t *testing.T) {
	at := NewAccessToken(APPID, SECRETE)
	a, err := at.Get()
	assert.Nil(t, err)
	assert.NotEmpty(t, a)
}
