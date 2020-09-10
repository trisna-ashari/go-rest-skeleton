package cmd_test

import (
	"go-rest-skeleton/infrastructure/persistence"
	"go-rest-skeleton/interfaces/cmd"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestNewCli(t *testing.T) {
	newCli := cmd.NewCli()

	var cliApp *cli.App
	assert.IsType(t, newCli, cliApp)
}

func TestNewCommand(t *testing.T) {
	var repositories *persistence.Repositories
	newCommand := cmd.NewCommand(repositories)

	var cliCommand []*cli.Command
	assert.IsType(t, cliCommand, newCommand)
}
