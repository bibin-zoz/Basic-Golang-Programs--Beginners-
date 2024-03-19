package main

import (
	"fmt"
)

func main() {

	ninja := []string{"ninja", "evil", "dead"}
	b := make(chan bool)

	for _, name := range ninja {
		go ninjaName(name, b)
	}
	fmt.Println(<-b)

}

func ninjaName(a string, status chan bool) {
	defer fmt.Println("Over", a)
	fmt.Println(a)
	status <- true
}

//WITHOUT CHANNELS..WE NEED TO SET SLEEP IN MAIN CODE...SO IT WILL WAIT
//UNTILL SECOND FUNC DONE EXECUTED..ELSE THE SECOND PROCESS NOT GONNA COMPLETE BCZZ PROGRAM CLOSES AFTER FUNC COMPLETES EXE
