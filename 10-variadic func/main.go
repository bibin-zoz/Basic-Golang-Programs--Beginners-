package main

import "fmt"

func sum(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total = total + val
	}
	fmt.Println(values[2])
	return total, "success"
}
func dis(values ...interface{}) {

	fmt.Println(values...)

}

func main() {

	a, status := sum(1, 2, 3, 45)
	fmt.Println(a)
	fmt.Println(status)
	dis(10, "hiii", true, 10.5)

}
