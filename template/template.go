package template

import (
	"bytes"
	"text/template"

	"github.com/JakobLybarger/leetcode-emailer/leetcode"
)

func GenerateTemplate() string {
	file := "./template.html"
	tmpl, err := template.New(file).ParseFiles(file)
	if err != nil {
		panic(err)
	}

	userStats := leetcode.GetLeetCodeStats()

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, userStats)
	if err != nil {
		panic(err)
	}

	return buf.String()
}
