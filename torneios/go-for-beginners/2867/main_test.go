package main

import (
	"fmt"
	"math"
	"testing"
)

func Test2867(t *testing.T) {
	cases := []struct {
		N float64
		M float64
		R int64
	}{
		{1, 1, 1},
		{2, 10, 4},
		{3, 9, 5},
		{100, 100, 201},
		{3, 3, 2},
		{2, 5, 2},
		{2, 8, 3},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("%.0f^%.0f", test.N, test.M), func(t *testing.T) {
			count := counter(math.Pow(test.N, test.M))
			if count != test.R {
				t.Errorf("Esperado %d, obtido %d", test.R, count)
			}
		})
	}
}
