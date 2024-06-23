package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "Your_slack_bot_token")
	os.Setenv("SLACK_APP_TOKEN", "Your_slack_app_token")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	channelID, timestamp, err := api.PostMessage(
		"C071L5GNZ9T",
		slack.MsgOptionText("Hello People", false),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Messafe sent successfully to channel %s at %s", channelID, timestamp)

}
// this send basic message to channel id=C071L5GNZ9T