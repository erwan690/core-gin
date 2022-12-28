package commands

import (
	"testing"

	"core-gin/lib"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/mock"
	"go.uber.org/fx"
)

type mockCommand struct {
	mock.Mock
}

func (m *mockCommand) Short() string {
	return "mock command"
}

func (m *mockCommand) Setup(cmd *cobra.Command) {}

func (m *mockCommand) Run() lib.CommandRunner {
	return func(fx.Lifecycle) {}
}

func TestWrapSubCommand(t *testing.T) {
	name := "test"
	cmd := &mockCommand{}
	opt := fx.Options()

	wrappedCmd := WrapSubCommand(name, cmd, opt)
	if wrappedCmd.Use != name {
		t.Errorf("expected wrapped command Use field to be %q, got %q", name, wrappedCmd.Use)
	}

	if wrappedCmd.Short != cmd.Short() {
		t.Errorf("expected wrapped command Short field to be %q, got %q", cmd.Short(), wrappedCmd.Short)
	}
}
