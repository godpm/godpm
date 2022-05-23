package rpc

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"

	"github.com/godpm/godpm/pkg/log"
	"github.com/lucas-clemente/quic-go/http3"
)

// NewQuicClient new quic client
func NewQuicClient(secret string, opt ...interface{}) Clienter {
	pool, err := x509.SystemCertPool()
	if err != nil {
		log.Error().Panic("get system cert pool failed", err)
		return nil
	}
	return &HTTPClient{
		Client: &http.Client{
			Transport: &HTTPTransport{
				secret: secret,
				tr: &http3.RoundTripper{
					TLSClientConfig: &tls.Config{
						RootCAs:            pool,
						InsecureSkipVerify: false,
					},
				},
			},
		},
	}
}
