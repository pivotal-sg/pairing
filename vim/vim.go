package vim

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

var (
	customPreConfig = "~/.config/nvim/custom_preconfig"
	customConfig    = "~/.config/nvim/custom_config"
)

func init() {
	var err error
	customPreConfig, err = filepath.Abs(customPreConfig)
	if err != nil {
		panic(err.Error())
	}
	customConfig, err = filepath.Abs(customConfig)
	if err != nil {
		panic(err.Error())
	}
}

func deleteIfSymlink(filename string) error {
	fileInfo, err := os.Lstat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if int(fileInfo.Mode())|int(os.ModeSymlink) != 0 {
		return os.Remove(filename)
	}
	return fmt.Errorf("'%s' is not a symbolic link", filename)
}

func linkCustomConfig(newPreConfig, newConfig string) error {
	for _, conf := range []string{customPreConfig, customConfig} {
		if err := deleteIfSymlink(conf); err != nil {
			return err
		}
	}
	if newPreConfig != "" {
		if err := os.Symlink(newConfig, customPreConfig); err != nil {
			return err
		}
	}
	if newConfig != "" {
		if err := os.Symlink(newConfig, customConfig); err != nil {
			return err
		}
	}
	return nil
}

var Command cli.Command = cli.Command{
	Name:  "relink",
	Usage: "Manually specify the directories you want as your preconfig and custom config.  Either are optional",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "pre-conf",
			Value: "",
			Usage: "Directory of the custom preconfig",
		},
		cli.StringFlag{
			Name:  "conf",
			Value: "",
			Usage: "Directory of the custom config",
		},
	},
	Action: func(c *cli.Context) error {
		preconf, conf := c.String("pre-conf"), c.String("conf")
		return linkCustomConfig(preconf, conf)
	},
}
