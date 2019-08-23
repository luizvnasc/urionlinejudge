package main

import "fmt"

func sum(a, b int64) int64 {
	return a + b
}

func main() {
	var a, b int64
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Printf("X = %d\n", sum(a, b))
}
