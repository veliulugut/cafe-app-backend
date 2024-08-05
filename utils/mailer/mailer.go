package mailer

import (
	"crypto/tls"
	"os"
	"strconv"

	"gopkg.in/mail.v2"
)

var (
	SMTP_MAIL_USERNAME string = os.Getenv("SMTP_MAIL_USERNAME")
	SMTP_MAIL_PASSWORD string = os.Getenv("SMTP_MAIL_PASSWORD")
	SMTP_HOST          string = os.Getenv("SMTP_HOST")
	SMTP_PORT          string = os.Getenv("SMTP_PORT")
)

func SendMail(To []string, Cc []string, Bcc []string, Subject string, HTMLBody string, attachmentFilePaths []string) error {
	m := mail.NewMessage()

	m.SetHeader("From", SMTP_MAIL_USERNAME)
	m.SetHeader("To", To...)
	m.SetHeader("Cc", Cc...)
	m.SetHeader("Bcc", Bcc...)
	m.SetBody("text/html", HTMLBody, mail.SetPartEncoding(mail.Unencoded))

	for _, attachmentFilePaths := range attachmentFilePaths {
		m.Attach(attachmentFilePaths)
	}

	port, err := strconv.Atoi(SMTP_PORT)
	if err != nil {
		return err
	}

	d := mail.NewDialer(
		SMTP_HOST,
		port,
		SMTP_MAIL_USERNAME,
		SMTP_MAIL_PASSWORD,
	)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
