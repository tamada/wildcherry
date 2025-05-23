package wildcherry

import (
	"bufio"
	"io"
	"strings"
)

type Result struct {
	Bytes uint64
	Chars uint64
	Words uint64
	Lines uint64
}

type CountingTargets struct {
	Bytes      bool
	Characters bool
	Words      bool
	Line       bool
}

type SourceList interface {
	IsDir() bool
	HasNext() bool
	Next() (Source, error)
}

type ReaderSource struct {
	source  io.ReadCloser
	scanner *bufio.Scanner
	scanned bool
}

func NewReaderSource(source io.ReadCloser) *ReaderSource {
	return &ReaderSource{
		source:  source,
		scanner: bufio.NewScanner(source),
		scanned: false,
	}
}

func (u *ReaderSource) HasNext() bool {
	u.scanned = true
	return u.scanner.Scan()
}

func (u *ReaderSource) Next() (string, error) {
	if !u.scanned {
		if !u.scanner.Scan() {
			return "", u.scanner.Err()
		}
	}
	u.scanned = false
	return u.scanner.Text(), u.scanner.Err()
}

func (u *ReaderSource) Close() error {
	return u.source.Close()
}

type Source interface {
	HasNext() bool
	Next() (string, error)
	Close() error
}

func NewResult() *Result {
	return &Result{
		Bytes: 0,
		Chars: 0,
		Words: 0,
		Lines: 0,
	}
}

func Count(source Source, ct *CountingTargets) (*Result, error) {
	r := NewResult()
	for source.HasNext() {
		line, err := source.Next()
		if err != nil {
			return nil, err
		}
		if ct.Bytes {
			r.Bytes += uint64(len(line))
		}
		if ct.Characters {
			r.Chars += uint64(len([]rune(line)))
		}
		if ct.Words {
			words := strings.Split(line, ",.:;!? \n\t ")
			r.Words += uint64(len(words))
		}
		if ct.Line {
			r.Lines++
		}
	}
	return r, nil
}
