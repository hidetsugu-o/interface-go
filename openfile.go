package main

import (
	"bytes"
	"io"
	"os"
)

type FileObject interface {
	Read(p []byte) (n int, err error)
	Close() error
}

func OpenFile(path string) (io.ReadCloser, error) {
	return os.Open(path)
}

func ReadAll(r io.Reader) ([]byte, error) {
	buf := bytes.Buffer{}
	_, err := buf.ReadFrom(r)
	return buf.Bytes(), err
}
