package functions

import (
	"fmt"
	"time"
)

func Init() {
	//x := 5
	//y := func() int {
	//	return x * 2
	//}()
	//
	//fmt.Println(y)

	c := make(chan int)

	go func() {
		fmt.Println("Hello World")
		time.Sleep(5 * time.Second)
		fmt.Println("Bye World")
		c <- 1
	}()
	<-c
}
