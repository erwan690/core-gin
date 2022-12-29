package infrastructure

import (
	"core-gin/lib"

	"github.com/slack-go/slack"
)

type ISlack interface {
	AlertSlack(message *MessageSlack) error
}

type Slack struct {
	*slack.Client
	env *lib.Env
}

func NewSlack(env *lib.Env) ISlack {
	slackClient := slack.New(env.SlackToken)

	return &Slack{
		slackClient,
		env,
	}
}

type MessageSlack struct {
	Message string
	Title   string
	Color   string
	Field   []slack.AttachmentField
}

func (s *Slack) AlertSlack(message *MessageSlack) error {
	// Create the Slack attachment that we will send to the channel
	message.Field = append(message.Field, slack.AttachmentField{
		Title: "PIC",
		Value: s.env.SlackMaintener,
	})
	attachment := slack.Attachment{
		Pretext: message.Title,
		Text:    message.Message,
		// Color Styles the Text, making it possible to have like Warnings etc.
		Color: message.Color,
		// Fields are Optional extra data!
		Fields: message.Field,
	}
	// PostMessage will send the message away.
	// First parameter is just the channelID, makes no sense to accept it
	_, _, err := s.PostMessage(
		s.env.SlackChannelID,
		// uncomment the item below to add a extra Header to the message, try it out :)
		// slack.MsgOptionText("New message from bot", false),
		slack.MsgOptionAttachments(attachment),
	)
	if err != nil {
		return err
	}
	return nil
}
