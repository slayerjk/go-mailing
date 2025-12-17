package mailing

import (
	"fmt"
	"net/smtp"
	"strings"
)

// Send plain text mail without auth(typically smtp:25).
func SendTextEmailWoAuth(mailHost string, mailPort int, mailFrom string, mailSubject string, body string, mailToList []string) error {
	// setting mail params
	smtpHostAndPort := fmt.Sprintf("%s:%d", mailHost, mailPort)

	// forming message
	msg := fmt.Sprintf("From: %s\r\n", mailFrom)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mailToList, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mailSubject)
	msg += fmt.Sprintf("\r\n%s\r\n", body)

	// Send the email
	err := smtp.SendMail(smtpHostAndPort, nil, mailFrom, mailToList, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}

// Send HTML mail without auth(typically smtp:25).
// subject will be like "<YOUR SUBJECT> - (<DATE"02.01.2006 15:04">)"
func SendHtmlEmailWoAuth(mailHost string, mailPort int, mailFrom string, mailSubject string, body string, mailToList []string) error {
	// setting mail params
	smtpHostAndPort := fmt.Sprintf("%s:%d", mailHost, mailPort)

	// forming message
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mailFrom)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mailToList, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mailSubject)
	msg += fmt.Sprintf("\r\n%s\r\n", body)

	// Send the email
	err := smtp.SendMail(smtpHostAndPort, nil, mailFrom, mailToList, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}
