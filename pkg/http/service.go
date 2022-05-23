package http

import (
	"fmt"
	"net/url"

	"github.com/godpm/godpm/pkg/rpc"
)

type Service struct {
	addr   string
	client rpc.Clienter
}

// NewService new service
func NewService(addr, secret string, httpClient bool) *Service {
	c := rpc.NewHTTPClient(secret)
	if !httpClient {
		c = rpc.NewQuicClient(secret)
	}
	return &Service{
		addr:   addr,
		client: c,
	}
}

// Stop stop process
func (svc *Service) Stop(name string) (err error) {
	url := fmt.Sprintf("%s/v1/stop/%s", svc.addr, name)
	err = svc.client.PutJSON(url, nil, nil)
	return err
}

// Status process status
func (svc *Service) Status(names []string) (resp []ProcStatus, err error) {
	query := url.Values{
		"names": names,
	}

	url := fmt.Sprintf("%s/v1/status?%s", svc.addr, query.Encode())
	err = svc.client.Get(url, &resp)
	return
}

// Restart restart process
func (svc *Service) Restart(name string) (err error) {
	url := fmt.Sprintf("%s/v1/restart/%s", svc.addr, name)
	err = svc.client.PutJSON(url, nil, nil)
	return err
}

// Restart restart process
func (svc *Service) Start(name string) (err error) {
	url := fmt.Sprintf("%s/v1/start/%s", svc.addr, name)
	err = svc.client.PutJSON(url, nil, nil)
	return err
}
