package set

import "fmt"

func ExampleHasDupeRuneBool() {
	fmt.Println(HasDupeRuneBool("숨바꼭질"))
	fmt.Println(HasDupeRuneBool("다시합시다"))
	// Output:
	// false
	// true
}

func ExampleHasDupeRuneStruct() {
	fmt.Println(HasDupeRuneStruct("숨바꼭질"))
	fmt.Println(HasDupeRuneStruct("다시합시다"))
	// Output:
	// false
	// true
}
