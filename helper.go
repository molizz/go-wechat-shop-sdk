package shop

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/dghubble/sling"
)

func buildURL(p, accessToken string) string {
	u, err := url.Parse(baseAPI)
	if err != nil {
		panic(err)
	}
	u.Path = p
	query := u.Query()
	query.Add("access_token", accessToken)
	u.RawQuery = query.Encode()
	return u.String()
}

func buildEmptyMap() map[interface{}]interface{} {
	return map[interface{}]interface{}{}
}

func POST(accessToken, path string, from interface{}, bind interface{}) (*http.Response, error) {
	if from == nil {
		from = buildEmptyMap()
	}
	return doHTTP(sling.New().Post, accessToken, path, from, bind)
}

func GET(accessToken, path string, bind interface{}) (*http.Response, error) {
	return doHTTP(sling.New().Get, accessToken, path, nil, bind)
}

type doFunc func(path string) *sling.Sling

func doHTTP(do doFunc, accessToken, path string, from interface{}, bind interface{}) (*http.Response, error) {
	return do(buildURL(accessToken, path)).
		BodyJSON(from).
		ReceiveSuccess(bind)
}

func Upload(accessToken, apiPath, uploadFilePath string, bind interface{}) (*http.Response, error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	do := sling.New().Post

	file, err := os.Open(uploadFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	part1, err := writer.CreateFormFile("media", filepath.Base(uploadFilePath))
	_, err = io.Copy(part1, file)
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := do(buildURL(accessToken, apiPath)).Request()
	if err != nil {
		return nil, err
	}
	req.Body = io.NopCloser(payload)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	c := &http.Client{
		Timeout: 60 * time.Second,
	}
	return c.Do(req)
}
