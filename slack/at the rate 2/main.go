package main

import (
	"context"
	"log"
	"os"

	"github.com/shomali11/slacker"
	"github.com/slack-go/slack"
)

func main() {
	// Set the Slack bot token as an environment variable
	os.Setenv("SLACK_BOT_TOKEN", "Your_slack_bot_token")
	os.Setenv("SLACK_APP_TOKEN", "Your_slack_pp_token")

	// Initialize the bot client with the bot token
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), "")

	// Handle interactive events
	bot.Interactive(func(ctx slacker.InteractiveBotContext, event *slack.InteractionCallback) error {
		// Handle interactive events
		if event.Type != slack.InteractionTypeBlockActions {
			return nil
		}

		if len(event.ActionCallback.BlockActions) != 1 {
			return nil
		}

		action := event.ActionCallback.BlockActions[0]
		if action.BlockID != "approval-msg" {
			return nil
		}

		var text string
		switch action.ActionID {
		case "approve":
			text = "Let's go ahead with this sales lead!"
		case "reject":
			text = "Let's look for better sales leads..."
		case "remindLater":
			text = "Remind me about this in 4 hours"
		default:
			text = "I don't understand your mood..."
		}

		// Post message to the channel using Slack API
		slackClient := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
		if _, _, err := slackClient.PostMessage(event.Channel.ID, slack.MsgOptionText(text, false)); err != nil {
			log.Println("Error posting message:", err)
		}

		return nil
	})

	// Handle commands
	bot.Command("sales_lead", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			happyBtn := slack.NewButtonBlockElement("approve", "true", slack.NewTextBlockObject("plain_text", "Approve", true, false))
			happyBtn.Style = "primary"
			sadBtn := slack.NewButtonBlockElement("reject", "false", slack.NewTextBlockObject("plain_text", "Reject", true, false))
			sadBtn.Style = "danger"
			remindLaterBtn := slack.NewButtonBlockElement("remindLater", "true", slack.NewTextBlockObject("plain_text", "Remind later", true, false))
			remindLaterBtn.Style = "primary"

			msgBody := "New lead: days = 5, reward = $10,000"
			err := response.Reply("", slacker.WithBlocks([]slack.Block{
				slack.NewSectionBlock(slack.NewTextBlockObject(slack.PlainTextType, msgBody, true, false), nil, nil),
				slack.NewActionBlock("approval-msg", happyBtn, sadBtn, remindLaterBtn),
			}))
			if err != nil {
				log.Println("Error replying to command:", err)
			}
		},
	})

	// Start listening for events
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
