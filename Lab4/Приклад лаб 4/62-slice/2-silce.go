package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	println(a)

	b := a[0:4:5]
	println(b)
	fmt.Printf("%v\n%v\n", a, b)

	b = append(b, 9)
	println(a)
	println(b)
	fmt.Printf("%v\n%v\n", a, b)
	/*
		b = append(b, 10)
		println(a)
		println(b)
		fmt.Printf("%v\n%v\n", a, b)
	*/
}
