package slicemethod

import "fmt"

func Example_slice_add() {
	src := []int{10, 20, 40, 50, 60}
	src = append(src[:3], src[2:]...)
	src[2] = 30

	fmt.Println(src)
	// Output: [10 20 30 40 50 60]
}

func Example_slice_add2() {
	src := []int{10, 20, 40, 50, 60}
	i := 0

	if i < len(src) {
		src = append(src[:3], src[2:]...)
		src[2] = 30
	} else {
		src = append(src, 30)
	}

	fmt.Println(src)
	// Output: [10 20 30 40 50 60]
}

func Example_slice_add3() {
	src := []int{10, 20, 40, 50, 60}

	src = append(src, 30)
	copy(src[3:], src[2:])
	src[2] = 30

	fmt.Println(src)
	// Output: [10 20 30 40 50 60]
}

func Example_slice_delete() {
	src := []int{10, 20, 30, 40, 50, 60}

	src = append(src[:2], src[3:]...)

	fmt.Println(src)
	// Output: [10 20 40 50 60]
}

// 순서가 바뀌어도 되는 경우 삭제
func Example_slice_delete2() {
	src := []int{10, 20, 30, 40, 50, 60}

	src[2] = src[len(src)-1]
	src = src[:len(src)-1]

	fmt.Println(src)
	// Output: [10 20 60 40 50]
}
