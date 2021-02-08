package go_server

import (
	"fmt"
	"log"
	"net"
)

func Run() {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
	for {
		fmt.Println("listen....")
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		go func() {
			defer conn.Close()
			conn.Write([]byte("from server"))
		}()
	}
}