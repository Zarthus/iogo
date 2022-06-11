package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestLoadMappings(t *testing.T) {
	dir, err := ioutil.ReadDir("./examples")
	if err != nil {
		return
	}

	mappings := loadMappings()
	found := 0

	for _, file := range dir {
		expected := strings.Replace(file.Name(), ".go", "", 1)

		for _, given := range mappings {
			if given.Name == expected {
				found++
				break
			}
		}
	}

	if len(mappings) != found {
		t.Fatalf("Number of mappings (%d) is not equal to files in ./example (%d)", len(mappings), found)
	}
}

func TestExec(t *testing.T) {
	value := false
	mappings := []mapping{
		{
			Name:     "foo",
			Runnable: func() { value = true },
		},
	}

	exec("foo", mappings)

	if value != true {
		t.Fatalf("Expected value to be true")
	}
}

func TestSelectables(t *testing.T) {
	res := selectables([]mapping{
		{
			Name:     "foo",
			Runnable: func() {},
		},
		{
			Name:     "bar",
			Runnable: func() {},
		},
	})

	if res[0] != "foo" {
		t.Fatalf("Expected res[0] to be foo, was %s", res[0])
	}
	if res[1] != "bar" {
		t.Fatalf("Expected res[1] to be bar, was %s", res[1])
	}
}
