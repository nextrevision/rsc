package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDirectory(t *testing.T) {
	dir, err := isDirectory("examples/")
	assert.Nil(t, err)
	assert.Equal(t, dir, true)

	dir, err = isDirectory("examples")
	assert.Nil(t, err)
	assert.Equal(t, dir, true)

	dir, err = isDirectory("foobar")
	assert.NotNil(t, err)
	assert.Equal(t, dir, false)
}
