/*
 *     Copyright (C) 2019  bn4t
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
	"os"
	"time"
)

var c *smtp.Client
var servername string

func initMailClient() {

	servername = os.Getenv("SMTP_ADDRESS")
	host, _, _ := net.SplitHostPort(servername)
	auth := smtp.PlainAuth("", os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"), host)

	// TLS config
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         host,
	}

	// Here is the key, we need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsConfig)
	if err != nil {
		log.Fatal(err)
	}

	// create the smtp client
	c, err = smtp.NewClient(conn, host)
	if err != nil {
		log.Fatal(err)
	}

	// Login using the provided credentials
	if err = c.Auth(auth); err != nil {
		log.Fatal(err)
	}
}

func sendCSPMail(domain string, documentUri string, referrer string, violatedDirective string, originalPolicy string, blockedUri string) {
	from := mail.Address{Name: "CSP-Handler", Address: os.Getenv("SENDER_EMAIL")}
	to := mail.Address{Name: "CSP-Handler", Address: os.Getenv("RECEIVER_EMAIL")}
	subj := "CSP violation for " + domain
	body := "A CSP violation occurred for " + domain + " at " + documentUri + "\n\n**Additional info:** \nReferrer: " + referrer + "\nViolated directive: " + violatedDirective +
		"\nOriginal policy: " + originalPolicy + "\nBlocked URI: " + blockedUri + "\n\nThis violation happened at " + time.Now().UTC().Format("2 Jan 2006 15:04:05") + " UTC."

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/plain; charset=\"utf-8\""
	headers["Content-Transfer-Encoding"] = "base64"

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	// To && From
	if err := c.Mail(from.Address); err != nil {
		log.Print(err)
	}

	if err := c.Rcpt(to.Address); err != nil {
		log.Print(err)
	}

	// send data
	w, err := c.Data()
	if err != nil {
		log.Print(err)
	}

	// write mail to writer
	_, err = w.Write([]byte(message))
	if err != nil {
		log.Print(err)
	}

	// close writer
	err = w.Close()
	if err != nil {
		log.Print(err)
	}
}
