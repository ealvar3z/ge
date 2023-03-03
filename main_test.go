package main

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeFileTypeDetect(t *testing.T) {
	geBuf := &gebuf{
		Path: "file.c",
	}

	geFileTypeDetect(geBuf)

	assert.Equal(t, CE_FILE_TYPE_C, geBuf.Type)
	assert.Equal(t, filepath.Ext(geBuf.Path), ".c")

	geBuf.Path = "file.py"
	geFileTypeDetect(geBuf)
	assert.Equal(t, CE_FILE_TYPE_PYTHON, geBuf.Type)
	assert.Equal(t, filepath.Ext(geBuf.Path), ".py")

	geBuf.Path = "file.txt"
	geFileTypeDetect(geBuf)
	assert.Equal(t, CE_FILE_TYPE_PLAIN, geBuf.Type)
	assert.Equal(t, filepath.Ext(geBuf.Path), ".txt")
}
