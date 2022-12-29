package infrastructure

import (
	"testing"

	"core-gin/lib"

	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
)

func TestAlertSlack(t *testing.T) {
	env := &lib.Env{
		SlackToken:     "test-token",
		SlackMaintener: "test-maintainer",
		SlackChannelID: "test-channel-id",
	}
	slackClient := NewSlack(env)

	message := &MessageSlack{
		Message: "test message",
		Title:   "test title",
		Color:   "test color",
		Field:   []slack.AttachmentField{},
	}
	err := slackClient.AlertSlack(message)
	assert.Error(t, err)
	assert.EqualError(t, err, "invalid_auth")
}
