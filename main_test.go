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
	if 0 != demo(flags{}) {
		resetStdInOut()
		t.Fail()
	}
	stdout.WriteString("one\n")
	if 0 != demo(flags{selectFlag: true}) {
		resetStdInOut()
		t.Fail()
	}

	stdout.WriteString("inp\n")
	if 0 != demo(flags{progressFlag: true}) {
		resetStdInOut()
		t.Fail()
	}

	if 0 != demo(flags{selectFlag: true, confirmFlag: true, helpFlag: true}) {
		resetStdInOut()
		t.Fail()
	}

	if 1 != demo(flags{selectFlag: true, confirmFlag: true}) {
		resetStdInOut()
		t.Fail()
	}

	resetStdInOut()
}
