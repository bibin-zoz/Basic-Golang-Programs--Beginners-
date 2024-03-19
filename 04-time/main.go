package main

import (
	"fmt"
	"time"
)

func main() {

	currentimee := time.Now()
	fmt.Printf("%T \n", currentimee)
	fmt.Println("tim=", currentimee)

	fmt.Println("formatted time", currentimee.Format("3:04:05 01-02-2006 Monday"))

}
