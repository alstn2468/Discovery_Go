package main

import "fmt"

func sol(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println("타잔이 10원짜리 팬티를 입고, 20원짜리 칼을 차고 노래를 한다.")
		fmt.Println("타잔이 20원짜리 팬티를 입고, 30원짜리 칼을 차고 노래를 한다.")
	}
}

func main() {
	sol(3)
}
