package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	c := make(chan int, 6)
	wg.Add(1)
	go count12("one", c, &wg)
	wg.Add(1)
	go count12("two", c, &wg)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	// fmt.Println(<-c)
	wg.Wait()
}

func count12(str string, c chan int, wg *sync.WaitGroup) {
	var i int
	for i = 0; i < 20; i++ {
		if i > 15 {
			c <- i
		}

		fmt.Println(str, " ", i)
	}
	// c <- i
	// fmt.Println(<-c)
	wg.Done()
}
