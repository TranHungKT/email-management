package utils

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"os"

	"gopkg.in/gomail.v2"
)

const (
	SENDER                     = "tranhung26122612@gmail.com"
	DEFAULT_TEMPLATE_SOURCE    = "static/templates/"
	DEFAULT_TEMPLATE_FILE_TYPE = ".html"
)

type Request struct {
	to      []string
	subject string
	body    string
}

var Templates *template.Template

func init() {
	Templates = template.Must(template.ParseGlob("static/templates/*.html"))
}

func NewRequest(to []string, subject string, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) SendEmailsByGoMail() {
	mail := gomail.NewMessage()
	mail.SetHeader("From", SENDER)
	mail.SetHeader("Subject", r.subject)
	Password := os.Getenv("EMAIL_PASSWORD")

	d := gomail.NewDialer("smtp.gmail.com", 465, SENDER, Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	mail.SetHeader("To", r.to...)
	mail.SetBody("text/html", r.body)
	if err := d.DialAndSend(mail); err != nil {
		fmt.Print(err)
		panic(err)
	}

}
func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	var buf bytes.Buffer
	if err := Templates.ExecuteTemplate(&buf, templateFileName+DEFAULT_TEMPLATE_FILE_TYPE, data); err != nil {
		panic(err)
	}
	r.body = buf.String()
	return nil
}

func SendEmails(toMails []string, subject string, templateName string, data interface{}) error {
	if len(toMails) == 0 {
		return nil
	}

	r := NewRequest(toMails, subject, "")
	if err := r.ParseTemplate(templateName, data); err != nil {
		fmt.Print("error", err)
		fmt.Print("\n")
		return err
	}
	r.SendEmailsByGoMail()
	return nil
}
