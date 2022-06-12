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

	mappingsExt := loadMappings() // remove psuedo "alL"
	mappings := mappingsExt[1:]
	finds := 0
	var found bool
	var missing []string

	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		expected := strings.Replace(file.Name(), ".go", "", 1)
		if expected[len(expected)-4:] == "test" {
			continue
		}

		found = false
		for _, given := range mappings {
			if given.Name == expected {
				found = true
				break
			}
		}

		if !found {
			missing = append(missing, expected)
		} else {
			finds++
		}
	}

	if len(missing) != 0 {
		t.Fatalf("Missing %d mappings: %s", len(missing), strings.Join(missing, ", "))
	}
	if len(mappings) != finds {
		t.Fatalf("Number of mappings (%d) is not equal to files in ./example (%d)", len(mappings), finds)
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
