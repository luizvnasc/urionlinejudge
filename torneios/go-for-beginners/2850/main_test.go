package main

import (
	"fmt"
	"testing"
)

func Test2850(t *testing.T) {
	cases := []struct {
		Input string
		R     string
	}{
		{"esquerda", "ingles"},
		{"direita", "frances"},
		{"nenhuma", "portugues"},
		{"as duas", "caiu"},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("Levantando a perna %q o papagaio fala %q", test.Input, test.R), func(t *testing.T) {
			lang := falaLoro(test.Input)
			if lang != test.R {
				t.Errorf("Esperava que o papagaio fala-se %q, mas ele falou %q", test.R, lang)
			}
		})
	}
}
