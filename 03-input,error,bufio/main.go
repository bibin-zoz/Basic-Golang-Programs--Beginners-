package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter")
	input, _ := reader.ReadString('\n')                          //stores error if any error occcures
	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //string hav space or line break so triming it out

	if err != nil {
		fmt.Println("error occcured:", err)
	} else {
		fmt.Println("number=", num)
	}

}
