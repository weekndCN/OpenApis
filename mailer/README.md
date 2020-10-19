# example 

```
package main

import (
	"crypto/tls"

	mailer "github.com/OpenApis/m/mailer"
)

func main() {
	s := mailer.SMTPServer{}

	// mail config
	From := "services@xxxx"
	To := []string{"weeknd.su@xxxx", "weeknd.stark@xxxx"}
	Cc := []string{"weeknd.stark@xxxx"}
	Bcc := []string{"weeknd.stark@xxxx"}
	Subject := "I am Weeknd Stark!!"
	Body := "Weeknd Stark \n\nGood editing!!"

	// host config
	s.Host = "smtp.exmail.qq.com"
	s.Port = 465
	s.UserName = "services@xxxx"
	s.PassWord = "xxxx"
	s.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         s.Host,
	}
	r := mailer.NewSMTPRequest(From, To, Cc, Bcc, Subject, Body)
	r.SendMail(&s)
}
```