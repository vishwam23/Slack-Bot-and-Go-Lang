package main

import (
	"context"
	"log"
	"os"

	"github.com/shomali11/slacker"
)

func main() {
	// Set the Slack bot token as an environment variable
	os.Setenv("SLACK_BOT_TOKEN", "Your_slack_bot_token")
	os.Setenv("SLACK_APP_TOKEN", "Your_slack_app_token")

	// Initialize the bot client with the bot token and app token
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	// Define slash commands and their handlers
	definitionNewEmployeeJoined := &slacker.CommandDefinition{
		Description: "Notify when a new employee joins the company",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("We need to create a form for new employee details")
		},
	}

	definitionNewProjectReceived := &slacker.CommandDefinition{
		Description: "Notify when a new project is received",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("We need to create a form for new project details")
		},
	}

	// Register the slash commands
	bot.Command("new_employee_joined", definitionNewEmployeeJoined)
	bot.Command("new_project_received", definitionNewProjectReceived)

	// Start listening for events
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
