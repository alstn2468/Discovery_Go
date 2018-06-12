package main

import "fmt"

func sol(n int) {
	for i := 1; i <= n; i++ {
		fmt.Println("타잔이 " + string(i*10) + "원짜리 팬티를 입고, " + string((i+1)*10) + "원짜리 칼을 차고 노래를 한다.")
	}
}

func main() {
	sol(3)
}
