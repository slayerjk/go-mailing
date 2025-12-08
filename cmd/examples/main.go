package main

import "fmt"

func main() {
	fmt.Print("test")
	/* // mailing example 'report'(read and send log file)
	report, err := os.ReadFile(logFile.Name())
	if err != nil {
		log.Fatal(err)
	}
	errM1 := mailing.SendPlainEmailWoAuth("mailing.json", "report", appName, report)
	if errM1 != nil {
		log.Printf("failed to send email:\n\t%v", errM1)
	} */

	/* // mailing example 'error'(just error text)
	newError := fmt.Errorf("custom error")
	errM2 := mailing.SendPlainEmailWoAuth("mailing.json", "error", appName, []byte(newError.Error()))
	if errM2 != nil {
		log.Printf("failed to send email:\n\t%v", errM2)
	} */

}
