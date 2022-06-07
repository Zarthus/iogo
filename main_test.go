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
	if demo(flags{}) != 0 {
		resetStdInOut()
		t.Fail()
	}

	stdout.WriteString("one\n")
	if demo(flags{selectFlag: true}) != 0 {
		resetStdInOut()
		t.Fail()
	}

	stdout.WriteString("inp\n")
	if demo(flags{progressFlag: true}) != 0 {
		resetStdInOut()
		t.Fail()
	}

	if demo(flags{selectFlag: true, confirmFlag: true, helpFlag: true}) != 0 {
		resetStdInOut()
		t.Fail()
	}

	if demo(flags{selectFlag: true, confirmFlag: true}) != 0 {
		resetStdInOut()
		t.Fail()
	}

	resetStdInOut()
}
