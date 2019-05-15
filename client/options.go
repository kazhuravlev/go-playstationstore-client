package client

import (
	"github.com/pkg/errors"
	"net/http"
)

var (
	ErrInvalidConfiguration = errors.New("invalid configuration")
)

type Option func(*Client) error

func WithHttpClient(httpClient *http.Client) Option {
	return func(c *Client) error {
		if httpClient == nil {
			return errors.Wrap(ErrInvalidConfiguration, "httpClient must be not nil")
		}

		c.http = httpClient
		return nil
	}
}
