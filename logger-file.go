package loggerfile

import (
	"io"
)

type options struct {
	nameFormat string
}

func New(nameFormat string) io.Writer {
	return options{
		nameFormat: nameFormat,
	}
}
