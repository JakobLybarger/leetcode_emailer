package model

type LeetCodeResponse struct {
	User        string
	Count       int          `json:"count"`
	Submissions []Submission `json:"submission"`
}
