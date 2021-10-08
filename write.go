package loggerfile

import (
	"os"
	"path/filepath"
	"time"
)

func (o options) Write(p []byte) (n int, err error) {
	path := time.Now().Format(o.nameFormat)
	folder := filepath.Dir(path)

	if _, err := os.Stat(folder); os.IsNotExist(err) {
		err := os.MkdirAll(folder, 0755)
		if err != nil {
			return 0, err
		}
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return f.Write(p)
}
