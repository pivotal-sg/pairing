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

// init the config directories.  This defaults to the neovim location
// because that is much nicer
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

// deleteIfSymlink will delete the pre and custom config dirs
// only if they exist, or if they are symlinks.
// It will error out if they are not symlinks, or on os errors, like
// a permission denied.
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

// linkCustomConfig is the meat of the customization of the
// vim config.  It assumes that the loading of the vim config
// goes something like :
//
//    runtime! custom_preconfig/*.vim
//    runtime! common_config/*.vim
//    runtime! custom_config/*.vim
//
// and it will link the pre and custom directories in.
// expected errors may os permission or file errors, as you would
// expect from  a file manipulation function
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

// RelinkCommand the urfave/cli command for manually relinking
var RelinkCommand cli.Command = cli.Command{
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
