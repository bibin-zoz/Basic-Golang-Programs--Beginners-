package main

import "fmt"

func main() {

	student("bibin", 10, 10, 20, 30)

}

func student(name string, id int, mark ...int) {

	var total int
	var count int
	fmt.Println("Student name:", name)
	fmt.Println("Student id:", id)
	for _, val := range mark {
		total = total + val
		count = count + 1
	}
	avg := total / count
	fmt.Println("avg mark=", avg)

}
