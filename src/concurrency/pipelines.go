package concurrency

import "fmt"

func Pipelines() {
	generator := make(chan int)
	doubles := make(chan int)

	go Generator(generator)
	go Double(generator, doubles)
	Print(doubles)
}

func Generator(c chan<- int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c)
}

func Double(in <-chan int, out chan<- int) {
	for value := range in {
		out <- 2 * value
	}
	close(out)
}

func Print(c <-chan int) {
	for value := range c {
		fmt.Println(value)
	}
}
