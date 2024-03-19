package main

import "fmt"

func addSlice(a []int) {
	b := len(a)
	a[b-1] = 0
	fmt.Println(" function call", a)

}

func main() {

	var a = [5]int{10, 21, 25, 66, 78}
	var b = []int{10, 50, 23, 15, 24}
	b = append(b, 5)

	fmt.Println(len(b), cap(b), b)

	//a = append(a,10) not possible in array
	b = append(b, 10) //possible....its extends if it hae capacity and else allocates new array
	c := a[1:5]

	fmt.Println(b, c)
	b = append(b[:3])
	index := 2
	c = append(c[:index], c[index+1:]...)
	fmt.Println(b, c)

	addSlice(b)
	fmt.Println("final", b, a)

	//slice remove

}
