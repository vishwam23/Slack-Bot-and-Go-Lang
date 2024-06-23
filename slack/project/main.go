package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "Your_slack_bot_token")
	os.Setenv("SLACK_APP_TOKEN", "Your_slack_app_token")

	args := os.Args[1:]
	fmt.Println(args)

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	preText := "Hello! your vishwam's build finished!"
	vishwamUrl := "*Build URL:* " + args[0]
	buildResult := "*" + args[1] + "*"
	buildNumber := "*" + args[2] + "*"
	jobName := "*" + args[3] + "*"

	if buildResult == "*SUCCESS*" {
		buildResult = buildResult + " :white_check_mark:"
	} else {
		buildResult = buildResult + " :x:"
	}

	dividerSection1 := slack.NewDividerBlock()
	vishwamBuildDetails := jobName + " #" + buildNumber + " - " + buildResult + "\n" + vishwamUrl
	preTextField := slack.NewTextBlockObject("mrkdwn", preText+"\n\n", false, false)
	vishwamBuildDetailsField := slack.NewTextBlockObject("mrkdwn", vishwamBuildDetails, false, false)

	vishwamBuildDetailsSection := slack.NewSectionBlock(vishwamBuildDetailsField, nil, nil)
	preTextSection := slack.NewSectionBlock(preTextField, nil, nil)

	msg := slack.MsgOptionBlocks(
		preTextSection,
		dividerSection1,
		vishwamBuildDetailsSection,
	)

	_, _, _, err := api.SendMessage(
		"C071L5GNZ9T",
		msg,
	)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}
// send this to terminal:go run main.go http://localhost SUCCESS 125 test-job
//       these are input given to       viswam url, buidresult, build-no, job-name    