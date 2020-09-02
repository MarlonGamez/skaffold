package ui

import (
	"io"
	"io/ioutil"
)

var (
	out io.Writer = ioutil.Discard
)

// SetOutput sets where UI functions will write to
func SetOutput(o io.Writer) {
	out = o
}
