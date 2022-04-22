package main

import (
	"bytes"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		n    int
		args []string
		want string
	}{
		{"add 2 to 4", 4, []string{}, "got 6"},
		{"add 2 to 44", 44, []string{}, "got 46"},
		{"add 2 to 49", 49, []string{}, "got 51"},
	}
	for _, tC := range testCases {
		out = new(bytes.Buffer)
		addTwo(tC.n)
		got := out.(*bytes.Buffer).String()
		if got != tC.want {
			t.Errorf("%s = %q, want %q", tC.desc, got, tC.want)
		}
	}
}
