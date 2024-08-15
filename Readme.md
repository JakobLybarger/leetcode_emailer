# LeetCode Accountability Tracker

This project is a Golang-based tool designed to help you and your friend stay accountable for daily LeetCode problem-solving. The program retrieves information about completed LeetCode problems using the [alfa-leetcode-api](https://github.com/alfaarghya/alfa-leetcode-api) and sends a daily email summarizing the number of problems solved, the specific problems, and the programming languages used.

## Features

- **Daily Email Summary**: Each day, an email is sent to both you and your friend, detailing:
  - The number of LeetCode problems solved.
  - The specific problems solved.
  - The programming languages used.

- **Golang Implementation**: The application is written in Go, leveraging its efficiency and simplicity.

- **Google Cloud Function Deployment**: The program is deployed as a Google Cloud Function, ensuring scalability and easy management.

## Installation

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16+)
- [Google Cloud SDK](https://cloud.google.com/sdk/docs/install)
- A valid Google Cloud account with billing enabled
- Access to the [alfa-leetcode-api](https://github.com/alfaarghya/alfa-leetcode-api)

