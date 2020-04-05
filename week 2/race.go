// The race condition is nothing but sharing a shared resource and the output
// cannot be predicted. Each time running the program will result in different outputs each time

package main

import "fmt"

var gVariable string

func main() {
	go print2(&gVariable)
	go print1(&gVariable)
	for {
		fmt.Print(gVariable)
	}
}

func print1(gVariable *string) {
	for {
		*gVariable = "u"
	}
}

func print2(gVariable *string) {
	for {
		*gVariable = "O"
	}
}
