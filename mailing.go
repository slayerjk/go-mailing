package mailing

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

// Send email, type=plain/html
func SendEmailWoAuth(mType string, mailHost string, mailPort int, mailFrom string, mailSubject string, body string, mailToList []string, attachPathes []string) error {
	// setting mail params
	smtpHostAndPort := fmt.Sprintf("%s:%d", mailHost, mailPort)
	var contentType string = ""

	switch mType {
	case "plain":
		contentType = "text/plain"
	case "html":
		contentType = "text/html"
	default:
		contentType = "text/plain"
	}

	var msg bytes.Buffer

	// forming message
	msg.WriteString(fmt.Sprintf("From: %s\r\n", mailFrom))
	msg.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(mailToList, ";")))
	msg.WriteString(fmt.Sprintf("Subject: %s\r\n", mailSubject))

	// setting boundary and content-type
	boundary := "my-boundary-779"
	msg.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", boundary))

	// writing body(text or html)
	msg.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
	msg.WriteString(fmt.Sprintf("Content-Type: %s; charset=utf-8\n", contentType))
	msg.WriteString(fmt.Sprintf("\r\n%s\r\n", body))

	// if there are attachment files
	if len(attachPathes) > 0 {

		for _, v := range attachPathes {
			msg.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
			msg.WriteString("Content-Transfer-Encoding: base64\n")
			msg.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\n", v))
			msg.WriteString(fmt.Sprintf("Content-ID: %s\r\n\r\n", v))

			data, err := os.ReadFile(v)
			if err != nil {
				return fmt.Errorf("failed to read attach file: %s, \n\t%v", data, err)
			}

			b := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
			base64.StdEncoding.Encode(b, data)
			msg.Write(b)

			msg.WriteString(fmt.Sprintf("\n--%s", boundary))
		}
		msg.WriteString("--")
	}

	// Send the email
	err := smtp.SendMail(smtpHostAndPort, nil, mailFrom, mailToList, msg.Bytes())
	if err != nil {
		return err
	}

	return nil
}
