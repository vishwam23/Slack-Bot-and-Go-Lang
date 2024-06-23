package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

var (
	botToken  = os.Getenv("SLACK_BOT_TOKEN")
	channelID = "C071L5GNZ9T" // Replace with your actual channel ID
)

func main() {
	// Ensure the bot token is set in the environment
	os.Setenv("SLACK_BOT_TOKEN", "Your_slack_bot_token")
	botToken = os.Getenv("SLACK_BOT_TOKEN")

	// Create a slice of slack.BlockElement
	buttons := []slack.BlockElement{
		slack.NewButtonBlockElement("approve_button", "approve", slack.NewTextBlockObject(slack.PlainTextType, "Approve", false, false)).WithStyle(slack.StylePrimary),
		slack.NewButtonBlockElement("reject_button", "reject", slack.NewTextBlockObject(slack.PlainTextType, "Reject", false, false)).WithStyle(slack.StyleDanger),
		slack.NewButtonBlockElement("asklater_button", "ask_later", slack.NewTextBlockObject(slack.PlainTextType, "Ask Later", false, false)),
	}

	message := slack.MsgOptionBlocks(
		slack.NewSectionBlock(slack.NewTextBlockObject(slack.PlainTextType, "Sales lead", false, false), nil, nil),
		slack.NewSectionBlock(slack.NewTextBlockObject(slack.MarkdownType, "A new sales lead is found.\nBudget - 10000", false, false), nil, nil),
		slack.NewActionBlock("actionblock", buttons...),
	)

	// Initialize the Slack API client
	api := slack.New(botToken)

	// Post the message to the specified channel
	channelID, timestamp, err := api.PostMessage(channelID, message)
	if err != nil {
		fmt.Printf("Failed to send message: %v\n", err)
		return
	}
	fmt.Printf("Message sent successfully to channel %s at %s\n", channelID, timestamp)
}
