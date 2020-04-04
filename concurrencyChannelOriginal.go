package main

import (
	"fmt"
	"strconv"
)

func main() {
	c := make(chan string)
	go count2("u", c)
	msg := <-c
	fmt.Println(msg)
}

func count2(str string, c chan string) {
	for i := 1; i <= 10; i++ {
		c <- strconv.Itoa(i)
		fmt.Println(i, " ", str)
	}

}
