package main

import (
	"crypto/tls"

	mailer "github.com/OpenApis/m/mailer"
)

func main() {

	var s mailer.SMTPServer
	// mail config
	From := "services@wukongbox.com"
	To := []string{"weeknd.su@wukongbox.com", "weeknd.stark@wukongbox.com"}
	Cc := []string{"weeknd.stark@wukongbox.com"}
	Bcc := []string{"weeknd.stark@wukongbox.com"}
	Subject := "I am Weeknd Stark!!"
	Body := "Harry Potter and threat to Israel\n\nGood editing!!"

	// host config
	s.Host = "smtp.exmail.qq.com"
	s.Port = 465
	s.Username = "services@wukongbox.com"
	s.Password = "WuKong@2017"
	s.TLSConfig = &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         s.Host,
	}
	r := mailer.NewSMTPRequest(From, To, Cc, Bcc, Subject, Body)
	r.SendMail(&s)
}
