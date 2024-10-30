package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func letsGetValue(chan1 chan int) {

	chan1 <- 12
}

func printWords(exit chan int) {
	for {
		select {
		case <-exit:
			return
		default:
			time.Sleep(500 * time.Microsecond)
			fmt.Println("Hello World")
		}
	}
}

func main() {
	b := 0
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			b = i
		}()
	}

	// chan1 := make(chan (int), 2)

	wg.Wait()
	// go letsGetValue(chan1)
	fmt.Println(b)

	// for i := 0; i < 100; i++ {
	// 	go func(ch chan int) {
	// 		ch <- i
	// 	}(chan1)
	// }

	// for va := range chan1 {
	// 	fmt.Println(va)
	// }

	exitChan := make(chan int)

	go printWords(exitChan)

	time.Sleep(5000 * time.Microsecond)

}
