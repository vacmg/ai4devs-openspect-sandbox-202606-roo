package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  int
	}{
		{"valido", "user@example.com\n", 0},
		{"valido_con_espacios", "  user@example.com  \n", 0},
		{"invalido", "no-arroba\n", 1},
		{"vacio", "", 1},
		{"invalido_dominio", "user@example\n", 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var stderr bytes.Buffer
			got := run(strings.NewReader(c.input), &stderr)
			if got != c.want {
				t.Errorf("run(%q) = %d, want %d", c.input, got, c.want)
			}
		})
	}
}

