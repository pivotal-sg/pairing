package main

import (
	"fmt"
	"os"

	"github.com/pivotal-sg/pairing/vim"
	"github.com/spf13/cobra"
)

var version string

var rootCommand = &cobra.Command{
	Use: "pairing",
}

func init() {
	rootCommand.AddCommand(vim.Command)
}

func main() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
