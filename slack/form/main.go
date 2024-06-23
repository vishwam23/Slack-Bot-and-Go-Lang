package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/slack-go/slack"
)

const (
	targetChannelID = "C071L5GNZ9T"
)

type formField struct {
	Title        string   `json:"title"`
	Type         string   `json:"type"`              // "text", "select", etc. (refer to Slack API docs)
	Name         string   `json:"name"`              // Used for identifying the field
	OptionValues []string `json:"options,omitempty"` // For select elements
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "Your_slack_bot_token")

	slackBotToken := os.Getenv("SLACK_BOT_TOKEN")
	if slackBotToken == "" {
		log.Fatal("SLACK_BOT_TOKEN environment variable is not set")
	}

	formElements := []formField{
		{
			Title: "What is your name?",
			Type:  "text",
			Name:  "name",
		},
		{
			Title:        "What is your favorite color?",
			Type:         "select",
			Name:         "color",
			OptionValues: []string{"Red", "Green", "Blue", "Other"},
		},
	}

	dialog := slack.Dialog{
		Title:      "My Custom Form",
		CallbackID: "form_submission",
		Elements:   convertFormFields(formElements),
	}

	api := slack.New(slackBotToken)

	err := api.OpenDialog(targetChannelID, dialog)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/slack/form_submission", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			log.Println("Error reading request body:", err)
			return
		}

		var payload slack.InteractionCallback
		if err := json.Unmarshal(body, &payload); err != nil {
			http.Error(w, "Error unmarshalling submission payload", http.StatusInternalServerError)
			log.Println("Error unmarshalling submission payload:", err)
			return
		}

		fmt.Printf("Form submitted by: %s\n", payload.User.Name)
		for _, action := range payload.ActionCallback.BlockActions {
			fmt.Printf("Field: %s, Value: %s\n", action.ActionID, action.SelectedOption.Value)
		}

		// Respond to Slack with an empty 200 OK response
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func convertFormFields(fields []formField) []slack.DialogElement {
	elements := make([]slack.DialogElement, 0, len(fields))
	for _, field := range fields {
		switch field.Type {
		case "text":
			elements = append(elements, &slack.DialogInput{
				Label: field.Title,
				Name:  field.Name,
				Type:  slack.InputTypeText,
			})
		case "select":
			options := make([]slack.DialogSelectOption, len(field.OptionValues))
			for i, option := range field.OptionValues {
				options[i] = slack.DialogSelectOption{
					Label: option,
					Value: option,
				}
			}
			elements = append(elements, &slack.DialogInputSelect{
				DialogInput: slack.DialogInput{
					Label: field.Title,
					Name:  field.Name,
					Type:  slack.InputTypeSelect,
				},
				Options: options,
			})
		}
	}
	return elements
}
