package commands_test

import (
	"core-gin/commands"
	"core-gin/lib"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var _ = Describe("WrapSubCommand", func() {
	It("returns a *cobra.Command with a valid Run field", func() {
		name := "test-command"
		cmd := &TestCommand{}
		opt := fx.Options()
		wrappedCmd := commands.WrapSubCommand(name, cmd, opt)
		Expect(wrappedCmd.Run).NotTo(BeNil())
		Expect(wrappedCmd.Run).To(BeAssignableToTypeOf(func(c *cobra.Command, args []string) {}))
	})
})

var _ = Describe("GetSubCommands", func() {
	It("returns a non-empty slice of *cobra.Command values", func() {
		opt := fx.Options()
		Expect(commands.GetSubCommands(opt)).NotTo(BeEmpty())
	})
})

type TestCommand struct{}

func (c *TestCommand) Short() string {
	return "test command"
}

func (c *TestCommand) Setup(cmd *cobra.Command) {}

func (c *TestCommand) Run() lib.CommandRunner {
	return func(lc fx.Lifecycle) {}
}
