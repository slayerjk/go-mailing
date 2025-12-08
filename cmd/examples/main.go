package main

import (
	"flag"
	"fmt"
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
	flag.Parse()

	// forming list of mailTo addresses
	mailToList := strings.Split(*mailTo, ",")

	err := mailing.SendPlainEmailWoAuth(*mailHost, *mailPort, *mailFrom, *mailSubject, *mailMsg, mailToList)
	if err != nil {
		fmt.Println(err)
	}
}
