package main

import (
	"fmt"
	"os"

	"github.com/pivotal-sg/pairing/vim"
	"github.com/urfave/cli"
)

var version string

func main() {
	app := cli.NewApp()
	app.Name = "pairing"
	app.Usage = "Manage your pairing needs.  Like low BO, or high caffeine"
	app.Version = version
	app.Commands = []cli.Command{
		vim.Command,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
	}
}
