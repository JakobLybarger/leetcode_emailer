package leetcode

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/JakobLybarger/leetcode-emailer/model"
)

func GetLeetCodeStats() []model.LeetCodeResponse {
	accounts := strings.Split(os.Getenv("accounts"), ",")

	done := make(chan model.LeetCodeResponse, len(accounts))
	for _, account := range accounts {
		go getUserStats(done, account)
	}

	var userStats []model.LeetCodeResponse
	for i := 0; i < len(accounts); i++ {
		userStats = append(userStats, <-done)
	}

	close(done)

	return userStats
}

func getUserStats(done chan model.LeetCodeResponse, account string) {

	baseUrl := "https://alfa-leetcode-api.onrender.com/"
	url := fmt.Sprintf("%s%s/acSubmission", baseUrl, account)
	req, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	var res model.LeetCodeResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		panic(err)
	}

	cst, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		panic(err)
	}

	now := time.Now().In(cst)
	test := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	midnight := test.Unix()

	var solvedToday []model.Submission
	for _, submission := range res.Submissions {
		submissionTime, err := strconv.ParseInt(submission.Timestamp, 10, 64)
		if err != nil {
			panic(err)
		}

		if submissionTime > midnight {
			solvedToday = append(solvedToday, submission)
		}
	}

	res.Count = len(solvedToday)
	res.User = account
	res.Submissions = solvedToday

	done <- res
}
