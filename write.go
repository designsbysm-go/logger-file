package loggerfile

import (
	"os"
	"path/filepath"
	"time"
)

func (o options) Write(p []byte) (n int, err error) {
	path := time.Now().Format(o.path)
	folder := filepath.Dir(path)

	err = o.fs.MkdirAll(folder, 0755)
	if err != nil {
		return 0, err
	}

	f, err := o.fs.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return f.Write(p)
}
