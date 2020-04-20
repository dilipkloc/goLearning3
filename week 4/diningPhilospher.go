package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex }

type Philos struct {
	leftCS, rightCS *ChopS
	name            int
	eaten           int
}

func (p Philos) isHungry() bool {
	return p.eaten < 4
}

func eat(results chan<- int, p *Philos, num int) {
	p.leftCS.Lock()
	p.rightCS.Lock()

	fmt.Printf("starting to eat %v \n", p.name)

	time.Sleep(time.Second)

	fmt.Printf("finishing eating %v \n", p.name)
	p.leftCS.Unlock()
	p.rightCS.Unlock()
	results <- num
}

func eatMore(philos []*Philos) bool {
	for i := 0; i < len(philos); i++ {
		if philos[i].isHungry() {
			return true
		}
	}
	return false
}

func getRandom(philos []*Philos, eating int) int {
	random := rand.Intn(len(philos))

	for i := 0; i < len(philos); i++ {
		if philos[random].isHungry() && random != eating {
			return random
		}
		random = (random + 1) % len(philos)
	}
	return random
}

func host(philos []*Philos) {

	eaters := make(chan int, 2)

	random1 := getRandom(philos, -1)
	go eat(eaters, philos[random1], random1)
	philos[random1].eaten = 1 + philos[random1].eaten

	for eatMore(philos) {
		random2 := getRandom(philos, random1)
		go eat(eaters, philos[random2], random2)
		philos[random2].eaten = 1 + philos[random2].eaten

		result := <-eaters
		if result == random1 {
			random1 = random2
		}
	}
	random1 = <-eaters
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philos, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philos{CSticks[i], CSticks[(i+1)%5], i + 1, 1}
	}

	wg.Add(1)
	go host(philos)
	wg.Wait()
}
