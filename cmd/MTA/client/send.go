package client

import (
	"log"
	"net/smtp"

	"github.com/BurntSushi/toml"
)

type Config struct {
	tranfser_server string
	smtp_headers    string
}

func SendSMTP(from string, rcpt string, data string) {
	var conf Config
	if _, err := toml.DecodeFile("./../config/config.toml", &conf); err != nil {
		log.Fatal(err)
	}

	err := smtp.SendMail(conf.tranfser_server, nil, from, []string{rcpt}, []byte(conf.smtp_headers+data))
	if err != nil {
		log.Fatal(err)
	}
}
