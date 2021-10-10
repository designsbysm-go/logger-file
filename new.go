package loggerfile

import (
	"io"

	"github.com/spf13/afero"
)

type options struct {
	fs   afero.Fs
	path string
}

func New(path string) io.Writer {
	if path == "" {
		return nil
	}

	return options{
		fs:   afero.NewOsFs(),
		path: path,
	}
}
