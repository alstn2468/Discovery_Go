package main

import "fmt"

func sol(n int) int {
	p, q := 0, 1

	for i := 0; i < n; i++ {
		p, q = q, p+q
	}

	return p
}

func main() {
	fmt.Println(sol(10))
}
