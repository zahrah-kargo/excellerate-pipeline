package main

import (
	"testing"
)

const (
	_helloWorld = "Hello world!"
)

func TestHelloWorld(t *testing.T) {
	res := HelloWorld()

	if res != _helloWorld {
		t.Fatalf("returned %s, expecting %s", res, _helloWorld)
	}
}
