package loggerfile

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/spf13/afero"
)

func TestShouldCreateDatedFolder(t *testing.T) {
	w := options{
		fs:   afero.NewMemMapFs(),
		path: "logs/2006/01/02/test.log",
	}
	p := []byte("test log")

	w.Write(p)

	path := time.Now().Format(w.path)
	folder := filepath.Dir(path)

	_, err := w.fs.Stat(folder)
	if err != nil {
		t.Errorf("error should be nil, got: %v", err)
	}

	_, err = w.fs.Stat(path)
	if err != nil {
		t.Errorf("error should be nil, got: %v", err)
	}
}

func TestShouldWriteToFile(t *testing.T) {
	w := options{
		fs:   afero.NewMemMapFs(),
		path: "logs/2006/01/02/test.log",
	}
	p := []byte("test log")

	n, err := w.Write(p)
	// fmt.Println(n, err)
	if err != nil {
		t.Errorf("should not have error, got: %v", err)
	} else if len(p) != n {
		t.Errorf("should write %d bytes, got: %d", len(p), n)
	}
}

func TestShouldFailMissingFolder(t *testing.T) {
	w := options{
		fs:   afero.NewReadOnlyFs(afero.NewMemMapFs()),
		path: "logs/2006/01/02/test.log",
	}
	p := []byte("test log")

	_, err := w.Write(p)
	if err == nil {
		t.Fatal("error should not be nil")
	}
}
