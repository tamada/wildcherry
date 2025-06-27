package wildcherry

import (
	"bufio"
	"errors"
	"io"
	"os"
)

type Source interface {
	Name() string
	Reader() (io.ReadCloser, error)
}

type Result struct {
	S     Source
	Bytes uint64
	Words uint64
	Lines uint64
	Err   error
}

type Option struct {
	RespectGitignore bool
	Recursive        bool
	T                *Targets
}

func NewOption() *Option {
	return &Option{
		RespectGitignore: true,
		Recursive:        true,
		T:                &Targets{Bytes: true, Words: true, Line: true},
	}
}

type Targets struct {
	Bytes bool
	Words bool
	Line  bool
}

func NewResult(s Source) *Result {
	return &Result{
		S:     s,
		Bytes: 0,
		Words: 0,
		Lines: 0,
	}
}

func (r *Result) Sum(other *Result) {
	if other == nil {
		return
	}
	r.Bytes += other.Bytes
	r.Words += other.Words
	r.Lines += other.Lines
	if r.Err == nil && other.Err != nil {
		r.Err = other.Err
	} else if r.Err != nil && other.Err != nil {
		r.Err = errors.Join(r.Err, other.Err)
	}
}

func NewStdinSource() Source {
	return &stdinSource{}
}

type stdinSource struct {
}

func (s *stdinSource) Name() string {
	return "<stdin>"
}

func (s *stdinSource) Reader() (io.ReadCloser, error) {
	return io.NopCloser(os.Stdin), nil
}

func CountRoutine(source Source, t *Targets, ch chan<- *Result) {
	go func() {
		r, err := Count(source, t)
		if err != nil {
			r = &Result{S: source, Bytes: 0, Words: 0, Lines: 0, Err: err}
		}
		ch <- r
	}()
}

func Count(source Source, t *Targets) (*Result, error) {
	r := NewResult(source)
	reader, err := source.Reader()
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	var prev byte = ' '
	in := bufio.NewReader(reader)
	for {
		d, err := in.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if t.Bytes {
			r.Bytes++
		}
		if t.Words {
			if isWhitespace(prev) && !isWhitespace(d) {
				r.Words++
			}
			prev = d
		}
		if t.Line {
			if d == '\n' {
				r.Lines++
			}
		}
	}
	return r, nil
}

func isWhitespace(b byte) bool {
	switch b {
	case ' ', '\t', '\n', '\r', '\v', '\f', ',', '.', ':', ';', '?', '!', '\x00':
		return true
	default:
		return false
	}
}
