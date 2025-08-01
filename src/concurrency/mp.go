package concurrency

import (
	"fmt"
	"time"
)

func MainMp() {
	c1 := make(chan int)
	c2 := make(chan int)
	d1 := 4 * time.Second
	d2 := 2 * time.Second

	go DoSomethingMp(d1, c1, 1)
	go DoSomethingMp(d2, c2, 2)

	for i := 0; i < 2; i++ {
		select {
		case channelMsg := <-c1:
			fmt.Println("Channel Message: ", channelMsg)
		case channelMsg2 := <-c2:
			fmt.Println("Channel two Message : ", channelMsg2)
		}
	}
}

func DoSomethingMp(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}
