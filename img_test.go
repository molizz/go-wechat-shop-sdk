package shop

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImg_Upload(t *testing.T) {
	at, err := NewAccessToken(APPID, SECRETE).Get()
	assert.Nil(t, err)
	assert.NotEmpty(t, at)

	dir := filepath.Join(os.Getenv("GOPATH"), "src/github.com/molizz/go-wechat-shop-sdk")
	img := filepath.Join(dir, "test/icon.jpg")

	result, err := NewImg(at).Upload(img)
	assert.Nil(t, err)
	assert.NotEmpty(t, result.ImgInfo)
}
