package main

type Writer interface {
	Write([]byte) (int, error)
}

type Document struct {
	Name    string
	Content []byte
}

type writer struct {
}

func (writer) Write([]byte) (int, error) { return 1, nil }

var _ Writer = (*writer)(nil)

func Save(f Writer, doc *Document) error {
	return nil
}
