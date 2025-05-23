package fs

import (
	"errors"
	"os"

	"github.com/tamada/wildcherry"
)

type FileSource struct {
	path   string
	source *wildcherry.ReaderSource
}

func New(path string) (*FileSource, error) {
	if path == "" {
		return nil, errors.New("given path is empty")
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &FileSource{
		path:   path,
		source: wildcherry.NewReaderSource(file),
	}, nil
}

func (f *FileSource) HasNext() bool {
	return f.source.HasNext()
}

func (f *FileSource) Next() (string, error) {
	return f.source.Next()
}

func (f *FileSource) Close() error {
	return f.source.Close()
}
