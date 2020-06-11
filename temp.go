package with

import (
	"io/ioutil"
	"os"
)

// Temp creates a new temporary file using ioutil.TempFile and calls function f.
// The function gets created file object as an argument. Upon exit Temp closes
// the temporary file.
func Temp(dir, pattern string, f func(f *os.File) error) (err error) {
	var t *os.File
	t, err = ioutil.TempFile(dir, pattern)
	if err != nil {
		return
	}
	defer func() {
		if errC := t.Close(); err == nil {
			err = errC
		}
	}()

	err = f(t)
	return
}

// TempDir creates a temporary directory using ioutil.TempDir and calls
// function f. Then it removes the directory with all its content upon f exit.
func TempDir(dir, pattern string, f func(name string) error) (err error) {
	var name string
	name, err = ioutil.TempDir(dir, pattern)
	if err != nil {
		return
	}
	defer func() {
		if errR := os.RemoveAll(name); err == nil {
			err = errR
		}
	}()

	err = f(name)
	return
}
