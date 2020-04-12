package main

import (
	"fmt"
	"sync"
)

type chops struct{ sync.Mutex }

type Philo struct {
	lefttcs, righttcs *chops
}

func (p Philo) eat(j int, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		p.lefttcs.Lock()
		p.righttcs.Lock()

		fmt.Println("Starting to eat ", j)

		p.righttcs.Unlock()
		p.lefttcs.Unlock()
		fmt.Println("Finishing eating ", j)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	CSticks := make([]*chops, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(chops)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]}
	}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat(i+1, &wg)
	}
	wg.Wait()
}
