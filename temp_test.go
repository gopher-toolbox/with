package with

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTempCall(t *testing.T) {
	err := Temp("", "", func(f *os.File) error {
		assert.Nil(t, f)
		return nil
	})
	assert.NoError(t, err)
}
