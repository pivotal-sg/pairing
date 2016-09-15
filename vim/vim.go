package vim

import (
	"github.com/pivotal-sg/pairing/vim/cmd"
	"github.com/spf13/cobra"
)

func init() {
	Command.AddCommand(cmd.Relink)
}

// Command is the `spf13/cobra` subcommand for vim, it holds the pairing specific
// subcommands, which are setup in the init function
var Command = &cobra.Command{
	Use:   "vim",
	Short: "vim is a collection of commands for managing paring configs for vim",
}
