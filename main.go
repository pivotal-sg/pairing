package main

import (
	"fmt"
	"os"

	"github.com/pivotal-sg/pairing/vim"
	_ "github.com/pivotal-sg/pairing/vim/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var version string

var rootCommand = &cobra.Command{
	Use: "pairing",
}

func init() {
	rootCommand.AddCommand(vim.Command)
}

func main() {
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.WatchConfig()
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
