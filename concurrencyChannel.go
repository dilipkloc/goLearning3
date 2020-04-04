package main

import (
	"fmt"
	"sync"
)

func main() {
	// make(c chan)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		count("dilip", &wg)
		// wg.Done()
	}()
	go count("dilip", &wg)
	go count1("kumar")
	wg.Wait()

}

func count1(name string) {
	for i := 1; i <= 10; i++ {
		fmt.Println(i, " ", name)
	}
}

func count(name string, wg *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {
		fmt.Println(i, " ", name)
	}
	wg.Done()
}
