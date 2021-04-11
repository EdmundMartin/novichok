package novichok

import (
	"context"
	"net/http"
	"time"
)

type ChessComClient struct {
	http.Client
	Transport *http.Transport
	Timeout   time.Duration
}

type ClientOptions struct {
	TLSHandshakeTimeout   time.Duration
	DisableKeepAlives     bool
	DisableCompression    bool
	MaxIdleConns          int
	MaxIdleConnsPerHost   int
	MaxConnsPerHost       int
	IdleConnTimeout       time.Duration
	ResponseHeaderTimeout time.Duration
	ExpectContinueTimeout time.Duration
}

func GetDefaultOptions(duration time.Duration) ClientOptions {
	return ClientOptions{
		TLSHandshakeTimeout:   duration,
		MaxIdleConns:          20,
		MaxIdleConnsPerHost:   20,
		MaxConnsPerHost:       50,
		IdleConnTimeout:       30 * time.Second,
		ResponseHeaderTimeout: duration,
	}
}

type ClientOption func(*ClientOptions) error

func NewChessComClient(timeout time.Duration, options ...ClientOption) (hr *ChessComClient) {

	opts := GetDefaultOptions(timeout)

	for _, opt := range options {
		if err := opt(&opts); err != nil {
			return nil
		}
	}

	transport := &http.Transport{
		TLSHandshakeTimeout:   opts.TLSHandshakeTimeout,
		DisableKeepAlives:     opts.DisableKeepAlives,
		DisableCompression:    opts.DisableCompression,
		MaxIdleConns:          opts.MaxIdleConns,
		MaxIdleConnsPerHost:   opts.MaxIdleConns,
		MaxConnsPerHost:       opts.MaxConnsPerHost,
		IdleConnTimeout:       opts.IdleConnTimeout,
		ResponseHeaderTimeout: opts.ResponseHeaderTimeout,
		ExpectContinueTimeout: opts.ExpectContinueTimeout,
	}

	client := http.Client{
		Timeout:   timeout,
		Transport: transport,
	}

	hr = &ChessComClient{
		Client:    client,
		Transport: transport,
		Timeout:   timeout,
	}

	return
}

func (client *ChessComClient) Do(ctx context.Context, req *http.Request) (*http.Response, context.CancelFunc, error) {

	c, cancel := context.WithTimeout(ctx, client.Timeout)

	resp, err := client.Client.Do(req.WithContext(c))

	return resp, cancel, err
}

func (client *ChessComClient) Stop() {
	client.Transport.CloseIdleConnections()
}
