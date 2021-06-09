package shop

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/dghubble/sling"
)

func buildAccessTokenURL(path, accessToken string) string {
	return buildURL(path, url.Values{"access_token": []string{accessToken}})
}

func buildURL(path string, values url.Values) string {
	u, err := url.Parse(baseAPI)
	if err != nil {
		panic(err)
	}
	u.Path = path
	u.RawQuery = values.Encode()
	return u.String()
}

func buildEmptyJSONBody() struct{} {
	return struct{}{}
}

func POST(accessToken, path string, from interface{}, bind interface{}) (*http.Response, error) {
	if from == nil {
		from = buildEmptyJSONBody()
	}
	u := buildAccessTokenURL(path, accessToken)

	return doHTTP(sling.New().Post, u, from, bind)
}

func GET(accessToken, path string, bind interface{}) (*http.Response, error) {
	u := buildAccessTokenURL(path, accessToken)
	return doHTTP(sling.New().Get, u, nil, bind)
}

type doFunc func(path string) *sling.Sling

func doHTTP(do doFunc, u string, from interface{}, bind interface{}) (*http.Response, error) {
	return do(u).
		BodyJSON(from).
		ReceiveSuccess(bind)
}

func Upload(accessToken, apiPath, uploadFilePath string, bind interface{}) error {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	file, err := os.Open(uploadFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	media, err := writer.CreateFormFile("media", filepath.Base(uploadFilePath))
	_, err = io.Copy(media, file)
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}

	c := &http.Client{
		Timeout: 60 * time.Second,
	}
	req, err := http.NewRequest("POST", buildAccessTokenURL(apiPath, accessToken), payload)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(bind)
}
