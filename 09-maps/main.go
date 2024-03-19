package main

import "fmt"

func main() {
	a := make(map[int]string)
	a[1] = "me"
	a[2] = "jgd"
	a[3] = "lknbdt"

	for key, value := range a {
		fmt.Printf("%d : %v \n", key, value)
	}
}
