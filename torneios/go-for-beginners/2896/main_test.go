package main

import (
	"fmt"
	"testing"
)

func Test2896(t *testing.T) {
	cases := []struct {
		N int
		K int
		R int
	}{
		{7, 4, 4},
		{4, 7, 4},
		{4000, 7, 574},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("N: %d, K: %d", test.N, test.K), func(t *testing.T) {
			r := max(test.N, test.K)
			if r != test.R {
				t.Errorf("Experado %d, obtido %d", test.R, r)
			}
		})
	}
}
