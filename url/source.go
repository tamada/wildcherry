package url

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	neturl "net/url"

	"github.com/tamada/wildcherry"
)

type URLSource struct {
	url *neturl.URL
}

func IsURL(rawURL string) bool {
	if rawURL == "" {
		return false
	}
	u, err := neturl.Parse(rawURL)
	if err != nil {
		return false
	}
	return u.Scheme == "http" || u.Scheme == "https"
}

func New(rawURL string) (wildcherry.Source, error) {
	if rawURL == "" {
		return nil, errors.New("given url is empty")
	}
	u, err := neturl.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	s := &URLSource{url: u}
	return s, nil
}

func (u *URLSource) Name() string {
	return u.url.String()
}

func (u *URLSource) Reader() (io.ReadCloser, error) {
	response, err := http.Get(u.url.String())
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch %s: %s", u.url.String(), response.Status)
	}
	return response.Body, nil
}
