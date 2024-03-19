package main

import "fmt"

const LoginToken string = "username" //First letter capital==public
func main() {

	var a = 20
	b := 10
	var c, d int = 10, 30
	sum := a + b + c + d
	fmt.Println("sum=", sum)
	fmt.Printf("%T \n", sum)

	var status bool
	fmt.Println("status:", status)
	fmt.Printf("%T \n", status)
	//fp shortcut to println

	var num uint8 = 255 // unint 8-64,int 8-64
	fmt.Println("value:", num)
	fmt.Printf("%T \n", num)

}
