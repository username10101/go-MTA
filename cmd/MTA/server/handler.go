package server

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

type SMTPData struct {
	RCPT string
	FROM string
	HELO string
	DATA string
}

const (
	SERVER_HOST string = "localhost"
	SERVER_PORT string = "25"
	SERVER_TYPE string = "tcp"
)

func Handle(connectionFunc func(string, SMTPData)) {
	listening, err := net.Listen(SERVER_TYPE, net.JoinHostPort(SERVER_HOST, SERVER_PORT))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer listening.Close()

	for {
		conn, err := listening.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go handleRequest(conn, connectionFunc)
	}
}

func handleRequest(conn net.Conn, connectionFunc func(string, SMTPData)) {
	buf := make([]byte, 1024)

	data := SMTPData{}

	for {
		_, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading:", err.Error())
			}
			break
		}

		switch true {
		case strings.HasPrefix(string(buf), "HELO"):
			data.HELO = string(buf)
			conn.Write([]byte("250 OK\r\n"))
		case strings.HasPrefix(string(buf), "MAIL"):
			data.FROM = string(buf)
			conn.Write([]byte("250 OK\r\n"))
		case strings.HasPrefix(string(buf), "RCPT"):
			data.RCPT = string(buf)
			conn.Write([]byte("250 OK\r\n"))
		case strings.HasPrefix(string(buf), "DATA"):
			data.DATA = string(buf)
			conn.Write([]byte("250 OK\r\n"))
			conn.Close()
		}
	}

	connectionFunc("smtp.mail.ru:25", data)
}
