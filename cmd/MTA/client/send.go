package client

import (
	"log"
	"net/smtp"
)

type SMTPData struct {
	RCPT string
	FROM string
	HELO string
	DATA string
}

func SendSMTP(host string, data SMTPData) {
	err := smtp.SendMail(host, nil, data.FROM, []string{data.RCPT}, []byte(data.DATA))
	if err != nil {
		log.Fatal(err)
	}
}
