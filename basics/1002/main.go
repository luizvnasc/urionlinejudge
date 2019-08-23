package main

import "fmt"

const pi = 3.14159

func area(r float64) float64 {
	return pi * (r * r)
}

func main() {
	var r float64
	fmt.Scanf("%f", &r)
	fmt.Printf("A=%.4f\n", area(r))
}
