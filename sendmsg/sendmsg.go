package sendmsg

import (
	"fmt"
	"net/smtp"
)

func SendMsg(msg string) {

	// Sender data.
	from := "yourmail"
	password := "yourpasswrd"

	// Receiver email address.
	to := []string{
		"dunawlad@gmail.com",
	}

	// smtp server  Google configuration.
	smtpHost := "smtp.gmail.com" 
	smtpPort := "587"

	// Message.
	eMsg := msg
	message := []byte(eMsg)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
	}
}
