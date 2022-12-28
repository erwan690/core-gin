package bootstrap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	// Call the NewApp function
	app := NewApp()

	// Assert that the returned value is an App instance with the correct command
	assert.Equal(t, rootCmd, app.Command)
}
