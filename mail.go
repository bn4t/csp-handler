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
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
	"time"
)

// sendCSPMail sends a CSP violation email to the in the config specified receiver
func sendCSPMail(domain string, documentUri string, referrer string, violatedDirective string, originalPolicy string, blockedUri string) error {
	from := mail.Address{Name: "CSP-Handler", Address: Config.SenderEmail}
	body := "A CSP violation occurred for " + domain + " at " + documentUri + "\n\n**Additional info:** \nReferrer: " + referrer + "\nViolated directive: " + violatedDirective +
		"\nOriginal policy: " + originalPolicy + "\nBlocked URI: " + blockedUri + "\n\nThis violation happened at " + time.Now().UTC().Format(time.RFC1123Z) + "."
	host, _, _ := net.SplitHostPort(Config.SmtpAddress)
	auth := smtp.PlainAuth("", Config.SmtpUsername, Config.SmtpPassword, host)

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = Config.ReceiverEmail
	headers["Subject"] = "CSP violation report for " + domain
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/plain; charset=\"utf-8\""

	// Setup message
	var msg string
	for k, v := range headers {
		msg += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	msg += "\r\n\r\n" + body
	err := smtp.SendMail(Config.SmtpAddress, auth, Config.SenderEmail, []string{Config.ReceiverEmail}, []byte(msg))
	if err != nil {
		log.Print("An error occurred while sending a csp violation mail:")
		log.Print(err)
	}
	return err
}
