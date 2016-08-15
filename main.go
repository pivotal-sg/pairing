package main

import (
	"fmt"
	"os"

	"github.com/pivotal-sg/pairing/vim"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "pairing"
	app.Usage = "Manage your pairing needs.  Like low BO, or high caffeine"
	app.Commands = []cli.Command{
		cli.Command{
			Name: "vim",
			Subcommands: []cli.Command{
				vim.Command,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
	}
}
