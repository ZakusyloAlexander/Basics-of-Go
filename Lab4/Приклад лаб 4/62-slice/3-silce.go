package main

import (
	"fmt"
)

func main() {

	var nil_slice1 []int
	nil_slice2 := *new([]int)

	var empty_slice1 = []int{}
	empty_slice2 := make([]int, 0)

	fmt.Println(nil_slice1 == nil, len(nil_slice1), cap(nil_slice1))
	fmt.Println(nil_slice2 == nil, len(nil_slice2), cap(nil_slice2))

	fmt.Println(empty_slice1 == nil, len(empty_slice1), cap(empty_slice1))
	fmt.Println(empty_slice2 == nil, len(empty_slice2), cap(empty_slice2))
}
