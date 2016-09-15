package vim

import (
	"github.com/pivotal-sg/pairing/vim/cmd"
	_ "github.com/pivotal-sg/pairing/vim/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	Command.AddCommand(cmd.Relink)
	viper.BindPFlag("vim.pre-config", cmd.Relink.Flags().Lookup("pre-conf"))
	viper.BindPFlag("vim.config", cmd.Relink.Flags().Lookup("conf"))
}

// Command is the `spf13/cobra` subcommand for vim, it holds the pairing specific
// subcommands, which are setup in the init function
var Command = &cobra.Command{
	Use:   "vim",
	Short: "vim is a collection of commands for managing paring configs for vim",
}
