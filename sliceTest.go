package main

import (
	"fmt"
)

func f_slice_item(s []int) {
	s[0] = 1
}

func f_slice_size(s []int) {
	v := []int{10, 20, 30}
	s = append(s, v...)
	fmt.Println("appernd in func ", s)
}

func f_slice_pointer_item(s *[]int) {
	(*s)[0] = 3
}

func f_slice_pointer_size(s *[]int) {
	v := []int{10, 20, 30}
	*s = append(*s, v...)
}

func f_iter(s *[]int) {
	for i, v := range *s {
		fmt.Println(i, " val= ", v)
	}
}

func main() {
	var s []int

	fmt.Println("slice by value")
	s = []int{0}
	f_slice_item(s)
	fmt.Println(s)

	s = []int{0}
	f_slice_size(s)
	fmt.Println(s)

	fmt.Println()
	fmt.Println("slice by pointer")
	s = []int{0}
	f_slice_pointer_item(&s)
	fmt.Println(s)

	s = []int{0}
	f_slice_pointer_size(&s)
	fmt.Println(s)

	fmt.Println()
	fmt.Println("iter")
	f_iter(&s)
}
