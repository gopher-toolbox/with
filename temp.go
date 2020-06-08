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
