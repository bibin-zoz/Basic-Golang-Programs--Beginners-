package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, 世界")

	//array
	/*nums := []int{10, 2, 3, 4, 5, 6, 7, 8}
	nums = append(nums, 9)
	fmt.Println(nums)
	for _, i := range nums {
		fmt.Println(i)
	}*/

	//maps
	/*stud := map[string]string{
		"A": "Apple",
		"B": "Banana",
		"C": "Car",
	}
	for key, value := range stud {
		fmt.Println(key, ":", value)
	}*/
	//strrings
	str := "Hello, 世界"
	for index, char := range str {
		fmt.Printf("Index: %d, Rune: %c\n", index, char)
	}
}
