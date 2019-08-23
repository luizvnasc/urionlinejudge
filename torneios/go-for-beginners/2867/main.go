package main

import (
	"fmt"
	"math"
)

func counter(n float64) (counter int64) {
	for n >= 1 {
		n = n / 10
		counter++
	}
	return
}

func main() {
	var c int
	var n, m int
	fmt.Scanf("%d", &c)
	for i := 0; i < c; i++ {
		fmt.Scanf("%d %d", &n, &m)
		fmt.Println(counter(math.Pow(float64(n), float64(m))))
	}
}
