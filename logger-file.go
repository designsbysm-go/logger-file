package loggerfile

import (
	"io"
)

type options struct {
	path string
}

func New(path string) io.Writer {
	return options{
		path: path,
	}
}
