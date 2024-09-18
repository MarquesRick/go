package app_test

import (
	. "get-hosts-cli/app"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
)

func TestProcess(t *testing.T) {
	app := Process()

	assert.Equal(t, "CLI Application", app.Name)
	assert.Equal(t, "Search IPs and server names on internet", app.Usage)

	assert.Len(t, app.Commands, 2)

	// Test IP command
	ipCmd := app.Commands[0]
	assert.Equal(t, "ip", ipCmd.Name)
	assert.Equal(t, "Search IPs on internet", ipCmd.Usage)
	assert.NotNil(t, ipCmd.Action)

	// Test servers command
	serversCmd := app.Commands[1]
	assert.Equal(t, "servers", serversCmd.Name)
	assert.Equal(t, "Search server name on internet", serversCmd.Usage)
	assert.NotNil(t, serversCmd.Action)

	// Test flags
	assert.Len(t, app.Commands[0].Flags, 1)
	hostFlag := app.Commands[0].Flags[0].(cli.StringFlag)
	assert.Equal(t, "host", hostFlag.Name)
	assert.Equal(t, "google.com.br", hostFlag.Value)
}
