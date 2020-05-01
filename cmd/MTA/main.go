package main

import (
	"./client"
	"./server"
)

func main() {
	server.Handle(client.SendSMTP)
}
