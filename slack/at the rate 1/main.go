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
	os.Setenv("SLACK_APP_TOKEN", "Your_slack_app_token")

	// Initialize the bot client with the bot token
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), "")

	// Handle interactive events
	bot.Interactive(func(ctx slacker.InteractiveBotContext, event *slack.InteractionCallback) error {
		if event.Type != slack.InteractionTypeBlockActions {
			return nil
		}

		if len(event.ActionCallback.BlockActions) != 1 {
			return nil
		}

		action := event.ActionCallback.BlockActions[0]
		if action.BlockID != "mood-block" {
			return nil
		}

		var text string
		switch action.ActionID {
		case "happy":
			text = "I'm happy to hear you are happy!"
		case "sad":
			text = "I'm sorry to hear you are sad."
		default:
			text = "I don't understand your mood..."
		}

		// Post message to the channel
		slackClient := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
		if _, _, err := slackClient.PostMessage(event.Channel.ID, slack.MsgOptionText(text, false)); err != nil {
			log.Println("Error posting message:", err)
		}

		// Acknowledge the interaction event
		ctx.AckInteraction(event)

		return nil
	})

	// Handle commands
	bot.Command("mood", &slacker.CommandDefinition{
		Description: "Check your mood",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			happyBtn := slack.NewButtonBlockElement("happy", "true", slack.NewTextBlockObject("plain_text", "Happy 🙂", true, false))
			happyBtn.Style = "primary"
			sadBtn := slack.NewButtonBlockElement("sad", "false", slack.NewTextBlockObject("plain_text", "Sad ☹️", true, false))
			sadBtn.Style = "danger"

			// Send a message with interactive buttons
			msgBody := "What is your mood today?"
			err := response.Reply("", slacker.WithBlocks([]slack.Block{
				slack.NewSectionBlock(slack.NewTextBlockObject(slack.PlainTextType, msgBody, true, false), nil, nil),
				slack.NewActionBlock("mood-block", happyBtn, sadBtn),
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
