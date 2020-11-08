package golistcmd

import (
	"testing"
)

func TestSay(t *testing.T) {
	cases := []struct {
		name string
		in string
		out string
	}{
		{"helloworld", "world", "Hello world"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := say(c.in)
			if c.out != r {
				t.Errorf("[want] %v\t[got] %v", c.out, r)
			}
		})
	}
}