package main

import (
	"testing"
	"fmt"
)

func TestGenerate(t *testing.T) {
	fmt.Print("blah")
	prompt := generate("")
	if prompt != "" {
		t.Fail()
	}
}

func TestGenerate2(t *testing.T) {
	fmt.Print("blah")
	prompt := generate("")
	if prompt != "" {
		t.Fail()
	}
}
