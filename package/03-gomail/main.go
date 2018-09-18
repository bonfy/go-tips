package main

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

func main() {
	d := gomail.NewDialer("smtp.zoho.com", 587, "useremail", "password")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", "useremail")
	m.SetHeader("To", "targetmail")
	m.SetAddressHeader("Cc", "ccmail", "Bonfy")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	// m.Attach("/home/Alex/lolcat.jpg")

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
