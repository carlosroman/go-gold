package processor

import "io"

type Store interface {
	Save([]string)
	Flush() io.Reader
}

func NewStores() []Store {
	return nil
}
