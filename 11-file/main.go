package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "go langs file handling"

	file, err := os.Create("./gotext.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length:", length)

	defer file.Close()
	readFile("./gotext.txt")

}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("fil data: ", string(databyte))

}
