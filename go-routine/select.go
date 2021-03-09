package go_routine

import (
	"fmt"
	"sync"
)

func SelectLab() {
	ch := make(chan int)
	go SelectBlock1(ch)
	go SelectBlock2(ch)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(ch chan int) {
		ch<-1
		wg.Done()
	}(ch)
	wg.Wait()
	fmt.Println("fin")
}

func SelectBlock1(ch chan int) {
	select {
	case <-ch:
		fmt.Println("select goroutine 1")
	}
}

func SelectBlock2(ch chan int) {
	select {
	case <-ch:
		fmt.Println("select goroutine 2")
	}
}
