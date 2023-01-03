package infrastructure_test

import (
	"core-gin/infrastructure"
	"core-gin/lib"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/slack-go/slack"
)

var _ = Describe("AlertSlack", func() {
	It("returns an error if the Slack API request fails", func() {
		env := &lib.Env{
			SlackToken:     "test-token",
			SlackMaintener: "test-maintainer",
			SlackChannelID: "test-channel-id",
		}
		slackClient := infrastructure.NewSlack(env)

		message := &infrastructure.MessageSlack{
			Message: "test message",
			Title:   "test title",
			Color:   "test color",
			Field:   []slack.AttachmentField{},
		}
		Expect(slackClient.AlertSlack(message)).To(MatchError("invalid_auth"))
	})
})
