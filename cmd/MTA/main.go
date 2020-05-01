package main

import "./server"

func main() {
	server.Handle(client.SendSMTP)
}
