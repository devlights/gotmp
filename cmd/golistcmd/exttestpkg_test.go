package golistcmd_test

import (
	"testing"

	"github.com/devlights/gotmp/cmd/golistcmd"
)

func TestSay(t *testing.T) {
	cases := []struct {
		name string
		in string
		out string
	}{
		{"ext helloworld", "world", "Hello world"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := golistcmd.Say(c.in)
			if c.out != r {
				t.Errorf("[want] %v\t[got] %v", c.out, r)
			}
		})
	}
}