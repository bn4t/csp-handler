package main

import (
	"github.com/BurntSushi/toml"
	"log"
)

type configuration struct {
	SenderEmail   string `toml:"sender_email"`
	ReceiverEmail string `toml:"receiver_email"`
	SmtpAddress   string `toml:"smtp_address"`
	SmtpUsername  string `toml:"smtp_username"`
	SmtpPassword  string `toml:"smtp_password"`
	RateLimit     int    `toml:"rate_limit"`
	BindTo        string `toml:"bind_to"`
	ConfigPath    string
}

var Config configuration

func readConfig() {
	_, err := toml.DecodeFile(Config.ConfigPath, &Config)
	if err != nil {
		log.Fatal(err)
	}
}
