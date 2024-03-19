package main

import (
	"fmt"
	"sync"
)

var count = 0

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i <= 1000; i++ {
		wg.Add(1)
		go sum(&m, &wg)
	}
	wg.Wait()
	fmt.Println(count)
}

func sum(m *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Lock()
	count = count + 1
	m.Unlock()
}
