package review

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func Main() {
	var x int
	x = 8
	y := 7

	fmt.Println(x)
	fmt.Println(y)

	val, err := strconv.ParseInt("asdas", 0, 64)

	if err != nil {
		fmt.Println("Error converting 7", err)
	} else {
		fmt.Println(val)
	}

	m := make(map[string]int)

	m["a"] = 6

	fmt.Println(m["a"])

	s := []int{1, 2, 3}

	for i, v := range s {
		fmt.Println("Index: ", i)
		fmt.Println("Value: ", v)
	}

	s = append(s, 16)

	for i, v := range s {
		fmt.Println("Index: ", i)
		fmt.Println("Value: ", v)
	}

	// gorutines with waitgroup

	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()
			doSomethingWorker(i)
		}()

	}

	wg.Wait()

	//with channels

	c := make(chan int)
	go doSomething(c, 1)
	<-c

	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}

func doSomethingWorker(nro int) {
	fmt.Println("doSomething: ", nro)
	time.Sleep(3 * time.Second)
	fmt.Println("Done ", nro)
}

func doSomething(c chan int, nro int) {
	fmt.Println("doSomething: ", nro)
	time.Sleep(3 * time.Second)
	fmt.Println("Done ", nro)

	c <- nro
}
