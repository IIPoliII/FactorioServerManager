package admin

import "io"

// R is a reference to the running boolean in main.
var R *bool

// P references the var Pipe in main
var P *io.WriteCloser

var Stopped bool
