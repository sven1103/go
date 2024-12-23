package main

import "fmt"

func Slice() {
	var slice = make([]int, 0, 5)
	fmt.Println("emp: ", slice, "len: ", len(slice), "cap: ", cap(slice))
	slice = append(slice, 1)
	fmt.Println("emp: ", slice, "len: ", len(slice), "cap: ", cap(slice))
	slice = append(slice, 2)
	fmt.Println("emp: ", slice, "len: ", len(slice), "cap: ", cap(slice))
	slice = append(slice, 3, 4, 5, 6, 7)
	fmt.Println("emp: ", slice, "len: ", len(slice), "cap: ", cap(slice))
}
