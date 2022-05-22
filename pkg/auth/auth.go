package auth

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"net/http"

	"github.com/valyala/bytebufferpool"
)

// SignRequest sign request
func SignRequest(req *http.Request, secret string) (signStr string, err error) {
	sk := []byte(secret)
	h := hmac.New(sha256.New, sk)
	u := req.URL
	data := req.Method + " " + u.Path
	if u.RawQuery != "" {
		data = data + "?" + u.RawQuery
	}

	ctType := req.Header.Get("Content-Type")
	if ctType != "" {
		_, _ = io.WriteString(h, "\nContent-Type: "+ctType)
	}

	_, _ = io.WriteString(h, "\n\n")

	if needSignBody(req, ctType) {
		s2, err1 := readRequestBody(req)
		if err1 != nil {
			err = err1
			return
		}
		_, _ = h.Write(s2.Bytes())
		bytebufferpool.Put(s2)
	}

	sign := h.Sum(nil)
	signStr = base64.URLEncoding.EncodeToString(sign)
	return
}

func needSignBody(req *http.Request, contentType string) bool {
	return req.ContentLength != 0 && req.Body != nil && contentType != "" && contentType != "application/octet-stream"
}

func readRequestBody(req *http.Request) (buf *bytebufferpool.ByteBuffer, err error) {
	buf = bytebufferpool.Get()
	_, err = io.Copy(buf, req.Body)
	if err != nil {
		return
	}

	req.Body.Close()

	req.Body = io.NopCloser(bytes.NewBuffer(buf.Bytes()))
	return
}
