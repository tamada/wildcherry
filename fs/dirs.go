package fs

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/tamada/wildcherry"
)

func NewFromDir(path string, opts *wildcherry.Option) ([]wildcherry.Source, error) {
	if path == "" {
		return nil, errors.New("given path is empty")
	}
	ignore := newIgnore(path, opts.RespectGitignore)
	if !opts.Recursive {
		files, err := listFiles(path, ignore)
		if err != nil {
			return nil, err
		}
		return newFromFiles(files)
	}
	return newFromDir(path, opts, ignore)
}

func newFromFiles(files []string) ([]wildcherry.Source, error) {
	var r []wildcherry.Source
	for _, file := range files {
		r = append(r, &FileSource{path: file})
	}
	return r, nil
}

func listFiles(base string, ig Ignore) ([]string, error) {
	files, err := os.ReadDir(base)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, file := range files {
		path := filepath.Join(base, file.Name())
		if file.Type().IsRegular() && !ig.IsIgnore(path) {
			result = append(result, path)
		}
	}
	return result, nil
}

func newFromDir(base string, opts *wildcherry.Option, ig Ignore) ([]wildcherry.Source, error) {
	var r []wildcherry.Source
	err := filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !ig.IsIgnore(path) {
			r = append(r, &FileSource{path: path})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return r, nil
}
