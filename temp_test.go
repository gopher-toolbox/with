package with

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTempCall(t *testing.T) {
	var name string

	assert.NoError(t,
		Temp("", "", func(f *os.File) error {
			if !assert.NotNil(t, f) {
				t.FailNow()
			}
			name = f.Name()
			t.Logf("created file %q", name)
			return nil
		}),
	)
	assert.NoError(t, os.Remove(name))
}

func TestTempError(t *testing.T) {
	var name string

	assert.Error(t,
		Temp("/the/directory/which/should/not/exist", "", func(f *os.File) error {
			name = f.Name()
			t.Logf("created file %q", name)
			return nil
		}),
	)
	if name != "" {
		assert.NoError(t, os.Remove(name))
	}
}

func TestTempDirCall(t *testing.T) {
	assert.NoError(t,
		TempDir("", "", func(name string) error {
			if !assert.NotEmpty(t, name) {
				t.FailNow()
			}
			t.Logf("created directory %q", name)
			return nil
		}),
	)
}

func TestTempDirError(t *testing.T) {
	assert.Error(t,
		TempDir("/the/directory/which/should/not/exist", "", func(name string) error {
			t.Logf("created directory %q", name)
			return nil
		}),
	)
}
