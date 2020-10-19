package mailer

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"strconv"
	"strings"
)

// ServerInfo return host:Port
func (s *SMTPServer) ServerInfo() string {
	return s.Host + ":" + strconv.Itoa(s.Port)
}

// NewSMTPRequest return a new mail request
func NewSMTPRequest(from string, to, bcc, cc []string, subject, body string) *SMTPRequest {
	return &SMTPRequest{
		from:    from,
		to:      to,
		bcc:     bcc,
		cc:      cc,
		subject: subject,
		body:    body,
	}
}

// SendMail send mail
func (r *SMTPRequest) SendMail(s *SMTPServer) {
	msgBody := r.processMsg()

	// PlainAuth("", name , password, host)
	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	// detech server reachable
	conn, err := tls.Dial("tcp", s.ServerInfo(), s.TLSConfig)
	if err != nil {
		log.Panic(err)
	}

	client, err := smtp.NewClient(conn, s.Host)
	if err != nil {
		log.Panic(err)
	}

	// Auth client
	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	if err = client.Mail(r.from); err != nil {
		log.Panic(err)
	}
	// BCC hanlder

	receivers := append(r.to, r.cc...)
	receivers = append(receivers, r.bcc...)

	for _, k := range receivers {
		if err = client.Rcpt(k); err != nil {
			log.Panic(err)
		}
	}

	// Data
	w, err := client.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = fmt.Fprintf(w, msgBody)
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	client.Quit()

	log.Println("Mail sent successfully")
}

func (r *SMTPRequest) processMsg() (header string) {
	header += fmt.Sprintf("From: %s\r\n", r.from) // From: sender

	if len(r.to) > 0 { // To: Receiver
		header += fmt.Sprintf("To: %s\r\n", strings.Join(r.to, ";"))
	}

	if len(r.cc) > 0 { // Cc: Receiver
		header += fmt.Sprintf("Cc: %s\r\n", strings.Join(r.cc, ";"))
	}

	header += fmt.Sprintf("Subject: %s\r\n", r.subject)                 // Subject: mail subject
	header += fmt.Sprintf("Content-Type: text/html; charset=UTF-8\r\n") // Content-type
	header += "\r\n" + r.body                                           // mail Content

	return
}
