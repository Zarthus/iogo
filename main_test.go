package main

import (
	"os"
	"testing"
)

func TestFunc(t *testing.T) {
	oldStdout := os.Stdout
	defer func() {
		os.Stdout = oldStdout
	}()

	os.Stdout = nil
	if 0 != demo([]string{"iogo", "--confirm"}) {
		t.Fail()
	}
	if 0 != demo([]string{"iogo", "--help"}) {
		t.Fail()
	}
}
