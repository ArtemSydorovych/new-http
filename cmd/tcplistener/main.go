package main

import (
	"fmt"
	"http/internal/request"
	"log"
	"net"
)

const port = ":42069"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening for TCP traffic: %s\n", err.Error())
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("error accepting TCP connection: %s\n", err.Error())
		}
		r, err := request.RequestFromReader(conn)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Request line:")
		fmt.Printf("- Method: %s\n", r.RequestLine.Method)
		fmt.Printf("- Target: %s\n", r.RequestLine.RequestTarget)
		fmt.Printf("- Version: %s\n", r.RequestLine.HttpVersion)
	}
}
