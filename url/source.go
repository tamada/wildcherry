package url

import (
	"errors"
	"net/http"
	neturl "net/url"

	"github.com/tamada/wildcherry"
)

type URLSource struct {
	url    *neturl.URL
	source *wildcherry.ReaderSource
}

func New(rawURL string) (*URLSource, error) {
	if rawURL == "" {
		return nil, errors.New("given url is empty")
	}
	u, err := neturl.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	response, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	s := &URLSource{
		url:    u,
		source: wildcherry.NewReaderSource(response.Body),
	}
	return s, nil
}

func (u *URLSource) HasNext() bool {
	return u.source.HasNext()
}

func (u *URLSource) Next() (string, error) {
	return u.source.Next()
}

func (u *URLSource) Close() error {
	return u.source.Close()
}
