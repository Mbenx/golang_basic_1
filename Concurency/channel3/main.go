package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// add // wait // done

	wg.Add(2)
	go sayHello("Heru", &wg)

	// wg.Add(1)
	go sayHello("Herlambang", &wg)

	wg.Wait()
	// time.Sleep(1 * time.Second)
}

func sayHello(text string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println("hi ", text)
	}

	wg.Done()
}
