package main

import (
	"fmt"
	"math"
	"testing"
)

func Test1002(t *testing.T) {
	cases := []struct {
		raio float64
		A    float64
	}{
		{2.00, 12.5664},
		{100.64, 31819.3103},
		{150.00, 70685.7750},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("Área do circulo com raio = %f", test.raio), func(t *testing.T) {
			a := area(test.raio)
			if !assertFloatEquals(t, a, test.A) {
				t.Errorf("esperado é A: %.4f, obtido A: %.4f", test.A, a)
			}
		})
	}
}

//Valores abaixo de 4 digitos são irrelevantes
const eps float64 = 0.0001

func assertFloatEquals(t *testing.T, a, b float64) bool {
	if math.Abs(a-b) < eps {
		return true
	}
	return false
}
