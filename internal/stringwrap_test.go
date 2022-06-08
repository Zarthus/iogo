package internal

import (
	"strings"
	"testing"
)

func TestWrap(t *testing.T) {
	cases := []struct {
		Input     string
		Output    string
		MaxLength uint
	}{
		{
			"foo",
			"foo",
			4,
		},
		{
			"foobarbaz",
			"foob\narba\nz",
			4,
		},
	}

	for i, tc := range cases {
		actual := strings.Join(Wrap(tc.Input, tc.MaxLength), "\n")

		if actual != tc.Output {
			t.Fatalf(
				"Case %d Input:\n\n`%s`\n\n"+
					"Expected Output:\n\n`%s`\n\n"+
					"Actual Output:\n\n`%s`",
				i,
				tc.Input,
				tc.Output,
				actual,
			)
		}
	}
}
