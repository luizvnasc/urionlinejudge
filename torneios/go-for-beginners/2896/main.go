package main

import "fmt"

func max(n, k int) int {
	return n%k + n/k
}

func main() {
	var t, n, k int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		fmt.Scanf("%d %d", &n, &k)
		fmt.Println(max(n, k))
	}
}
