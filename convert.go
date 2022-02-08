package main

type File interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
	Seek(int64, int) (int64, error)
	Close() error
}

type Document struct {
	Name string
}

func Save(f File, doc *Document) error {
	return nil
}
