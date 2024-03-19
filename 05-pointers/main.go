package main

import "fmt"

func main() {
	var ptr *int

	ptr = new(int) //memory allocated
	*ptr = 10
	fmt.Println(*ptr + 10)

	fmt.Println(*ptr)

}
