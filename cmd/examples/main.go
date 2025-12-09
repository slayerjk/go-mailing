package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	mailing "github.com/slayerjk/go-mailing"
)

func main() {
	// smtpHost string, smtpPort string, mailFrom string, subject string, mailToList []string, msg []byte
	mailHost := flag.String("mh", "mail.example.com", "mail host address")
	mailPort := flag.Int("mp", 25, "mail server port")
	mailFrom := flag.String("mf", "no-reply@example.com", "mail from address")
	mailSubject := flag.String("ms", "TEST SUBJECT", "mail subject")
	mailMsg := flag.String("msg", "TEST MESSAGE", "message text")
	mailTo := flag.String("mt", "user1@example.com", "mail adresses separated by coma")
	mailType := flag.String("type", "plain", "'plain' or 'html'")
	flag.Parse()

	// forming list of mailTo addresses
	mailToList := strings.Split(*mailTo, ",")

	switch *mailType {
	case "plain":
		// sending plain text
		err := mailing.SendTextEmailWoAuth(*mailHost, *mailPort, *mailFrom, *mailSubject, *mailMsg, mailToList)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("DONE")
	case "html":
		// sending html
		body := fmt.Sprintf("<html><body><h1>%s</h1></body></html>", *mailMsg)
		err := mailing.SendHtmlEmailWoAuth(*mailHost, *mailPort, *mailFrom, *mailSubject, body, mailToList)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("DONE")
	default:
		fmt.Println("wrong message type('type' flag)")
		os.Exit(1)
	}
}
