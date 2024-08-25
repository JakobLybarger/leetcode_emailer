package template

import (
	"bytes"
	"embed"
	"text/template"

	"github.com/JakobLybarger/leetcode-emailer/leetcode"
)

//go:embed template.html
var templateFS embed.FS

func GenerateTemplate() string {
	tmpl, err := template.ParseFS(templateFS, "template.html")
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
