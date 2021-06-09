package shop

import (
	"net/url"
	"time"

	"github.com/dghubble/sling"
)

var _ AccessTokenGetter = (*AccessToken)(nil)

type AccessToken struct {
	appid  string
	secret string

	accessToken       string
	accessTokenExpire time.Time // token到期时间
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

func (a *AccessToken) getAccessTokenFromCache() string {
	nowUnix := time.Now().Unix()
	expiredUnix := a.accessTokenExpire.Unix() - 10 // 减少10秒主要是为了避免误差，

	if len(a.accessToken) == 0 {
		return ""
	}
	if nowUnix > expiredUnix {
		return ""
	}
	return a.accessToken
}

func (a *AccessToken) Get() (string, error) {
	at := a.getAccessTokenFromCache()
	if len(at) > 0 {
		return at, nil
	}

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
