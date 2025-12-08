package mailing

import (
	"fmt"
	"net/smtp"
	"strings"
	"time"
)

// type MailData struct {
// 	Host          string   `json:"host"`
// 	Port          string   `json:"port"`
// 	AuthUser      string   `json:"auth_user"`
// 	AuthPass      string   `json:"auth_pass"`
// 	FromAddr      string   `json:"from_addr"`
// 	ToAddrErrors  []string `json:"to_addr_errors"`
// 	ToAddrReports []string `json:"to_addr_reports"`
// }

// // read json mailing data
// func readMailingData(dataFile string) (MailData, error) {
// 	var result MailData

// 	// open file to read
// 	data, err := os.ReadFile(dataFile)
// 	if err != nil {
// 		return result, fmt.Errorf("failed to read mailing data file:\n\t%v", err)
// 	}

// 	// read file content
// 	errU := json.Unmarshal(data, &result)
// 	if errU != nil {
// 		return result, fmt.Errorf("failed to unmarshall mailing data:\n\t%v", errU)
// 	}

// 	return result, nil
// }

// Send plain text mail without auth(typically smtp:25).
// subject will be like "<YOUR SUBJECT> - (<DATE"02.01.2006 15:04">)"
func SendPlainEmailWoAuth(mailHost string, mailPort int, mailFrom string, mailSubject string, msg string, mailToList []string) error {
	// setting date string for Subject
	curDate := time.Now().Format("02.01.2006 15:04")
	subjectEnriched := fmt.Sprintf("%s - (%v)\n", mailSubject, curDate)

	// setting mail params
	smtpHostAndPort := fmt.Sprintf("%s:%d", mailHost, mailPort)

	// set "TO:"" header - must be comma separated values string
	toHeader := strings.Join(mailToList, ",")

	// setting message body
	message := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s>\n\n%v", mailFrom, toHeader, subjectEnriched, msg)

	// Send the email
	errS := smtp.SendMail(smtpHostAndPort, nil, mailFrom, mailToList, []byte(message))
	if errS != nil {
		return fmt.Errorf("error in SendMail func:\n\t%v", errS)
	}

	return nil
}
