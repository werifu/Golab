package go_routine

import (
	"fmt"
	"time"
)

func ChannelWrite(ch chan int) {
	ch<-1
	fmt.Println("write 1 in channel")
	ch<-2
	fmt.Println("write 2 in channel")
}


func Produce(ch chan int) {
	for i:=0;i < 5; i++ {
		time.Sleep(time.Second*1)
		ch<-i
		fmt.Printf("produce %d\n", i)
	}
}

func Consume(ch chan int) {
	for {
		select {
		case i:=<-ch:
			fmt.Printf("consume %d\n", i)
		}
	}
}