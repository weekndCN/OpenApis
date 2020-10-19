package mailer

import "crypto/tls"

// SMTPRequest SMTP send request config
type SMTPRequest struct {
	from    string
	to      []string
	bcc     []string
	cc      []string
	subject string
	body    string
}

// SMTPServer smtp server info
type SMTPServer struct {
	Username  string
	Password  string
	Host      string
	Port      int
	TLSConfig *tls.Config
}

type versionMail struct {
	Title   string
	Date    string
	Version string
}

type welcomeMail struct {
	Name string
	Date string
}
