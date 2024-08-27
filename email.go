package email

import (
	"context"
	"os"
	"strings"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/JakobLybarger/leetcode-emailer/template"
	"github.com/cloudevents/sdk-go/v2/event"
	"gopkg.in/gomail.v2"
)

func init() {
	functions.CloudEvent("LeetCodeEmail", email)
}

func email(ctx context.Context, e event.Event) error {
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

	return nil
}
