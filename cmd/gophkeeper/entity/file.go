package entity

import (
	"io"
	"strings"
	"time"
)

type FileInfo struct {
	UpdateTime time.Time
}

type TextReadCloser struct {
	reader io.Reader
}

func NewTextReadCloser(text string) *TextReadCloser {
	return &TextReadCloser{
		reader: strings.NewReader(text),
	}
}

func (trc *TextReadCloser) Read(p []byte) (n int, err error) {
	return trc.reader.Read(p)
}

func (trc *TextReadCloser) Close() error {
	return nil
}
