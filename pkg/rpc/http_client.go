package client

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/godpm/godpm/pkg/auth"
	"github.com/godpm/godpm/pkg/erro"
)

// HTTPClient client
type HTTPClient struct {
	*http.Client
}

// HTTPTransport http transport
type HTTPTransport struct {
	secret string
	tr     http.RoundTripper
}

// RoundTrip implement http.RoundTripper
func (htr *HTTPTransport) RoundTrip(req *http.Request) (*http.Response, error) {

	sign, err := auth.SignRequest(req, htr.secret)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "godpm "+sign)

	return htr.tr.RoundTrip(req)
}

// NewHTTPClient new http client with secret
func NewHTTPClient(secret string) Clienter {
	return &HTTPClient{
		Client: &http.Client{Transport: &HTTPTransport{secret: secret}},
	}
}

// Get implement clienter
func (c *HTTPClient) Get(url string, resp interface{}) (err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	return c.do(req, resp)
}

func (c *HTTPClient) PostForm(url string, body io.Reader, contentType string, resp interface{}) (err error) {
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", contentType)

	return c.do(req, resp)
}

// PostJSON post with json
func (c *HTTPClient) PostJSON(url string, req, resp interface{}) (err error) {
	return c.doJSON(http.MethodPost, url, req, resp)
}

// PutJSON put with json
func (c *HTTPClient) PutJSON(url string, req, resp interface{}) (err error) {
	return c.doJSON(http.MethodPut, url, req, resp)
}

// Delete delete
func (c *HTTPClient) Delete(url string, resp interface{}) (err error) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return
	}

	return c.do(req, resp)
}

func (c *HTTPClient) doJSON(method string, url string, req, resp interface{}) (err error) {
	body, err := json.Marshal(req)
	if err != nil {
		return
	}

	buf := bytes.NewBuffer(body)

	httpReq, err := http.NewRequest(method, url, buf)
	if err != nil {
		return
	}

	httpReq.Header.Set("Content-Type", "application/json; charset=utf-8")

	return c.do(httpReq, resp)
}

func (c *HTTPClient) do(req *http.Request, ret interface{}) (err error) {
	resp, err := c.Client.Do(req)
	if err != nil {
		return
	}

	defer func() {
		_, _ = io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}()

	if resp.StatusCode/100 == 2 {
		if ret != nil && resp.ContentLength > 0 {
			err = json.NewDecoder(resp.Body).Decode(ret)
			return
		}
	}

	e := &erro.Error{}
	if resp.ContentLength > 0 {
		err = json.NewDecoder(resp.Body).Decode(e)
		if err == nil {
			err = e
		}
	}

	return
}
