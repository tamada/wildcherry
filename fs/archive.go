package fs

import (
	"errors"

	"github.com/tamada/wildcherry"
)

func NewFromArchive(path string, opts *wildcherry.Option) ([]wildcherry.Source, error) {
	// This function should handle the logic to read from an archive file.
	// For now, we will return an error indicating that this functionality is not implemented.
	return nil, errors.New("not implemented yet")
}
