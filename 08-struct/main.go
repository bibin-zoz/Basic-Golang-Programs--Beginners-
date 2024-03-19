package main

import "fmt"

type User struct {
	Name string
	id   int
}

func (p User) display() {
	p = User{"bib", 10}
	fmt.Println(p.Name, p.id)

}

func main() {

	var a User

	a.display()

}
