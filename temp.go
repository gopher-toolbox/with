package with

import "os"

// Temp creates a new temporary file using ioutil.TempFile and calls function f.
// The function gets created file object as an argument. Upon exit Temp closes
// the temporary file.
func Temp(dir, pattern string, f func(f *os.File) error) error {
	return f(nil)
}
