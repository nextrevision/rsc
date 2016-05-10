package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDirectory(t *testing.T) {
	dir, err := IsDirectory("../examples/")
	assert.Nil(t, err)
	assert.Equal(t, dir, true)

	dir, err = IsDirectory("../examples")
	assert.Nil(t, err)
	assert.Equal(t, dir, true)

	dir, err = IsDirectory("foobar")
	assert.NotNil(t, err)
	assert.Equal(t, dir, false)
}
