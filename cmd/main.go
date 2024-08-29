package main

import (
	"os"
	"strings"

	"github.com/JakobLybarger/leetcode-emailer/template"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func main() {
	godotenv.Load()

	m := gomail.NewMessage()

	sender := os.Getenv("sender")
	m.SetHeader("From", sender)

	recipients := strings.Split(os.Getenv("recipients"), ",")
	m.SetHeader("To", recipients...)

	m.SetHeader("Subject", "LeetCode Accountability Tracker")

	body := template.GenerateTemplate()
	m.SetBody("text/html", body)

	password := os.Getenv("password")
	d := gomail.NewDialer("smtp.gmail.com", 587, sender, password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
