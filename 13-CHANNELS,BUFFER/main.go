package main

import "fmt"

func main() {

	status := make(chan string, 1) //channel have memory space
	status <- "hiii"
	fmt.Println(<-status) //channel passes to the same methode
	message := make(chan string)
	go chanRun(message)
	for {
		value, open := <-message
		if !open {
			break
		}
		fmt.Println(value)

	}

}
func chanRun(a chan string) {
	for i := 0; i < 3; i++ {
		a <- fmt.Sprint("channel value", i)
	}
	close(a)
}
