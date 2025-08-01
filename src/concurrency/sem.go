package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Sem() {
	c := make(chan int, 5)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)
		go doSomethingSem(i, &wg, c)
	}
	wg.Wait()
}

func doSomethingSem(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Println("Hello World from Platzi!!")
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Printf("Id %d finished\n", i)
	<-c
}
