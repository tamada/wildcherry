package wildcherry_test

import (
	"testing"

	"github.com/tamada/wildcherry"
	"github.com/tamada/wildcherry/fs"
)

func TestCount(t *testing.T) {
	td := []struct {
		id     string
		source string
		want   *wildcherry.Result
	}{
		{"test1", "testdata/LICENSE", &wildcherry.Result{nil, 1071, 169, 21, nil}},
	}

	opts := wildcherry.NewOption()
	for _, data := range td {
		source, _ := fs.New(data.source, opts)
		got, err := wildcherry.Count(source[0], opts.T)
		if err != nil {
			t.Errorf("Count(%s) returned error: %v", data.id, err)
			continue
		}
		if got == nil {
			t.Errorf("Count(%s) returned nil result", data.id)
			continue
		}
		if got.Lines != data.want.Lines || got.Words != data.want.Words || got.Bytes != data.want.Bytes {
			t.Errorf("Count(%s) = %v, want %v", data.id, got, data.want)
		}
	}
}
