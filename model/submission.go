package model

type Submission struct {
	Title         string `json:"title"`
	TitleSlug     string `json:"titleSlug"`
	Timestamp     string `json:"timestamp"`
	StatusDisplay string `json:"statusDisplay"`
	Language      string `json:"lang"`
}
