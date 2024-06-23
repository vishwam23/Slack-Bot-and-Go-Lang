package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shomali11/slacker"
)

func printComandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()

	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "Your_slack_bot_token")
	os.Setenv("SLACK_APP_TOKEN", "Your_slack_app_token")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printComandEvents(bot.CommandEvents())

	bot.Command("ping", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("pong")
		},
	})
	bot.Command("hi", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("HI wellcome, I'm here to help you out")
		},
	})
	bot.Command("vishwam", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("HI vishwam,is this anything i can help u out")
		},
	})

	bot.Command("kunal", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("HI kunal,is this anything i can help u out")
		},
	})
	bot.Command("amit", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("HI amit,is this anything i can help u out")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
