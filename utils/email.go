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
	SENDER = "tranhung26122612@gmail.com"
)

type Request struct {
	to      []string
	subject string
	body    string
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

	mail.SetHeader("To", "ngocanhtranthikt2002@gmail.com")
	mail.SetBody("text/html", r.body)
	if err := d.DialAndSend(mail); err != nil {
		fmt.Print(err)
		panic(err)
	}

}
func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)

	if err = t.Execute(buf, data); err != nil {
		return err
	}

	r.body = buf.String()
	return nil
}

const DEFAULT_TEMPLATE_SOURCE = "/Users/TranHung/Documents/email-management/templates/"
const DEFAULT_TEMPLATE_FILE_TYPE = ".html"

func SendEmails(toMails []string, subject string, templateName string, data interface{}) error {
	if len(toMails) == 0 {
		return nil
	}

	r := NewRequest(toMails, subject, "")
	if err := r.ParseTemplate(DEFAULT_TEMPLATE_SOURCE+templateName+DEFAULT_TEMPLATE_FILE_TYPE, data); err != nil {
		fmt.Print("error", err)
		fmt.Print("\n")
		return err
	}
	r.SendEmailsByGoMail()
	return nil

}
