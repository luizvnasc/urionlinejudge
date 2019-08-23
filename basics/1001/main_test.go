package main

import (
	"fmt"
	"testing"
)

func Test1001(t *testing.T) {
	cases := []struct {
		A int64
		B int64
		X int64
	}{
		{A: 10, B: 9, X: 19},
		{A: -10, B: 4, X: -6},
		{A: 15, B: -7, X: 8},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d + %d = %d", test.A, test.B, test.X), func(t *testing.T) {
			x := sum(test.A, test.B)
			if x != test.X {
				t.Errorf("esperado é X: %d, obtido X: %d", test.X, x)
			}
		})
	}
}
