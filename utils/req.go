package utils

import (
	"crypto/tls"
	"github.com/go-resty/resty/v2"
	"net"
	"net/http"
	"time"
)

func Requests(url, body, proxy string, header map[string]string, timeout int) (*resty.Response, error) {
	transport := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		Proxy:              http.ProxyFromEnvironment,
		DisableCompression: true,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	}

	client := resty.New().
		SetTimeout(time.Duration(timeout) * time.Second).
		SetTransport(transport)
	if proxy != "" {
		client.SetProxy(proxy)
	}

	r := client.R()
	if header != nil {
		r.SetHeaders(header)
	}

	if body != "" {
		r.SetBody(body)
		return r.Post(url)
	}
	return r.Get(url)
}
