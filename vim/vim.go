package vim

import "github.com/urfave/cli"

// Command is the `urfave/cli` subcommand for vim
var Command cli.Command = cli.Command{
	Name: "vim",
	Subcommands: []cli.Command{
		RelinkCommand,
	},
}
