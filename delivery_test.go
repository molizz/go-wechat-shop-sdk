package shop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelivery_GetCompanyList(t *testing.T) {
	at, err := NewAccessToken(APPID, SECRETE).Get()
	assert.Nil(t, err)

	d := NewDelivery(at)
	got, err := d.GetCompanyList()
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(got.CompanyList), 1)
}
