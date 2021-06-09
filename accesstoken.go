package shop

import (
	"net/url"

	"github.com/dghubble/sling"
)

var _ AccessTokenGetter = (*AccessToken)(nil)

type AccessToken struct {
	appid  string
	secret string
}

func NewAccessToken(appid, secret string) *AccessToken {
	return &AccessToken{
		appid:  appid,
		secret: secret,
	}
}

type AccessTokenResult struct {
	Result
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

func (a *AccessToken) Get() (string, error) {
	u := buildURL("cgi-bin/token", url.Values{
		"grant_type": []string{"client_credential"},
		"appid":      []string{a.appid},
		"secret":     []string{a.secret},
	})

	result := &AccessTokenResult{}
	_, err := doHTTP(sling.New().Get, u, nil, result)
	if err != nil {
		return "", err
	}
	if !result.OK() {
		return "", result
	}
	return result.AccessToken, nil
}
