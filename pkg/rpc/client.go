package rpc

import (
	"io"
)

// Clienter client interface
type Clienter interface {
	// Get handle json response
	Get(url string, resp interface{}) (err error)

	PostForm(url string, body io.Reader, contentType string, resp interface{}) (err error)
	PostJSON(url string, req, resp interface{}) (err error)
	PutJSON(url string, req, resp interface{}) (err error)
	Delete(url string, resp interface{}) (err error)
}
