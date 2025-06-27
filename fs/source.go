package fs

import (
	"errors"
	"io"
	"os"
	"strings"

	"github.com/tamada/wildcherry"
)

type FileSource struct {
	path string
}

func (f *FileSource) Name() string {
	return f.path
}
func (f *FileSource) Reader() (io.ReadCloser, error) {
	return os.Open(f.path)
}

func New(path string, opts *wildcherry.Option) ([]wildcherry.Source, error) {
	if path == "" {
		return nil, errors.New("given path is empty")
	}
	if IsDir(path) && opts.Recursive {
		return NewFromDir(path, opts)
	} else if IsArchive(path) {
		return NewFromArchive(path, opts)
	} else if ExistFile(path) {
		return []wildcherry.Source{&FileSource{path: path}}, nil
	}
	return nil, errors.New("unsupported source type: " + path)
}

type sourceList struct {
	base string
}

func ExistFile(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && stat.Mode().IsRegular()
}

func IsDir(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

func IsArchive(path string) bool {
	if ExistFile(path) {
		return strings.HasPrefix(path, ".tar") ||
			strings.HasPrefix(path, ".zip") ||
			strings.HasPrefix(path, ".tar.gz") ||
			strings.HasPrefix(path, ".tar.bz2")
	}
	return false
}
