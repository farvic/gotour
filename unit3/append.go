package main

import "fmt"

func main0() {
	var s []int
	printSlice0(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice0(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice0(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice0(s)
}

func printSlice0(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
