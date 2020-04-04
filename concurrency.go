package main

import (
	"fmt"
	"time"
)

func main() {
	go count("u")
	go count1("o")
	fmt.Scanln()
}

func count(name string) {
	for i := 1; true; i++ {
		fmt.Print(name)
		time.Sleep(time.Millisecond * 500)
	}
}

func count1(name string) {
	for i := 1; true; i++ {
		fmt.Print(name)
		time.Sleep(time.Millisecond * 500)
	}
}
