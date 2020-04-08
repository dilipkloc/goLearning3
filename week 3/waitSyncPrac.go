package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var s []int
	var s1 int
	for {
		fmt.Println("Enter the values for the array, enter -1 to stop entering values")
		fmt.Scan(&s1)
		if s1 == -1 {
			break
		}
		s = append(s, s1)
	}
	fmt.Print("Unsorted elements: ")
	fmt.Println(s)
	fmt.Print("\n\n\n")
	length := len(s) / 4
	wg.Add(1)
	mySort(s, 0, length-1, &wg)
	wg.Add(1)
	mySort(s, length, length*2-1, &wg)
	wg.Add(1)
	mySort(s, length*2, length*3-1, &wg)
	wg.Add(1)
	mySort(s, length*3, len(s)-1, &wg)
	wg.Wait()
	fmt.Print("\n\n\nSorted Elements:")
	sort.Ints(s)
	fmt.Println(s)
}

func mySort(i []int, start int, end int, wg *sync.WaitGroup) {
	sorted := false
	for !sorted {
		sorted = true
		for localstart := start; localstart < end; localstart++ {
			if i[localstart] > i[localstart+1] {
				i[localstart], i[localstart+1] = i[localstart+1], i[localstart]
				sorted = false
			}
		}
	}
	fmt.Println("one part of the array")
	for j := start; j <= end; j++ {
		fmt.Print(i[j], " ")
	}
	fmt.Println()
	wg.Done()
}
