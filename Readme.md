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

- [Go](https://golang.org/doc/install) (version 1.21+)
- [Google Cloud SDK](https://cloud.google.com/sdk/docs/install)
- A valid Google Cloud account with billing enabled
- Access to the [alfa-leetcode-api](https://github.com/alfaarghya/alfa-leetcode-api)

### Google Cloud Setup

1. **Create a Pub/Sub Topic:**

You'll need a Pub/Sub topic to trigger the Cloud Function. Create on with the following command:

```bash
gcloud pubsub topics create scheduled-email
```

2. **Set up a Cloud Scheduler Job:**

The Cloud Scheduler will trigger the Pub/Sub topic at your desired time. Use the following command to create a scheduler job:

```bash
gcloud scheduler jobs create pubsub leetcode-email-job \
--schedule="0 8 * * *" \
--time-zone="America/Chicago" \
--topic=scheduled-email \
--message-body="{}"
```
This example schedules the job to run daily at 8 AM CST

3. **Deploy your function using the following command:**

```bash
gcloud functions deploy leetcode-emailer \
--gen2 \
--runtime=go121 \
--region=us-central1 \
--source=. \
--entry-point=LeetCodeEmail \
--trigger-topic=scheduled-email \
--env-vars-file .env.yaml
```

### Example `.env.yaml` File

```yaml
sender: "your-email@example.com"
recipients: "friend1@example.com,friend2@example.com"
accounts: "YourLeetCodeUsername,FriendsLeetCodeUsername"
password: "fake-password"
```

