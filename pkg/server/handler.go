package server

import (
	"fmt"
	"net"
	"os"
)

const (
	SERVER_HOST string = "localhost"
	SERVER_PORT string = "2222"
	SERVER_TYPE string = "tcp"
)

func Handle() {
	listening, err := net.Listen(SERVER_TYPE, net.JoinHostPort(SERVER_HOST, SERVER_PORT))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer listening.Close()
}
