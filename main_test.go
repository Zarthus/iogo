package main

import (
	"os"
	"testing"
)

func TestFunc(t *testing.T) {
	oldStdin, oldStdout := os.Stdin, os.Stdout
	stdin, stdout, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stdin, os.Stdout = stdin, nil
	resetStdInOut := func() {
		os.Stdin, os.Stdout = oldStdin, oldStdout
	}

	stdout.WriteString("hello\n")
	if !demo(flags{}) {
		resetStdInOut()
		t.Fail()
	}
	stdout.WriteString("one\n")
	if !demo(flags{selectFlag: true}) {
		resetStdInOut()
		t.Fail()
	}
	if demo(flags{selectFlag: true, confirmFlag: true}) {
		resetStdInOut()
		t.Fail()
	}
	if !demo(flags{selectFlag: true, confirmFlag: true, helpFlag: true}) {
		resetStdInOut()
		t.Fail()
	}
	resetStdInOut()
}
