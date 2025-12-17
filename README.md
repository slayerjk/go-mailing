# go-mailing
Go - va mailing helper functions for standart "net/smtp" package

Only SMTP without auth for now!

<h2>SendTextEmailWoAuth</h2>

```
func SendTextEmailWoAuth(mailHost string, mailPort int, mailFrom string, mailSubject string, body string, mailToList []string) error
```
Send plain text mail without auth(typically smtp:25).

<h2>SendHtmlEmailWoAuth</h2>

```
func SendHtmlEmailWoAuth(mailHost string, mailPort int, mailFrom string, mailSubject string, body string, mailToList []string) error
```
Send plain text mail without auth(typically smtp:25).

<h2>cmd/examples/main.go<h2>

Testing functions.

Flags
```
    smtpHost string, smtpPort string, mailFrom string, subject string, mailToList []string, msg []byte
	mailHost := flag.String("mh", "mail.example.com", "mail host address")
	mailPort := flag.Int("mp", 25, "mail server port")
	mailFrom := flag.String("mf", "no-reply@example.com", "mail from address")
	mailSubject := flag.String("ms", "TEST SUBJECT", "mail subject")
	mailMsg := flag.String("msg", "TEST MESSAGE", "message text")
	mailTo := flag.String("mt", "user1@example.com", "mail adresses separated by coma")
	mailType := flag.String("type", "plain", "'plain' or 'html'")
```
