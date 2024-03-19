package main

import "fmt"

func main() {

	var a = new([]int)     // allocates and returns its address
	var b = make([]int, 0) //allocate and initialize
	fmt.Println(a)
	fmt.Println(b)
}
