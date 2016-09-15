package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	preConfigFlag, configFlag string
)

// init the config directories.  This defaults to the neovim location
// because that is much nicer, it also sets up your flags
func init() {
	Relink.Flags().StringVar(&preConfigFlag, "pre-conf", "", "Directory of the custom preconfig")
	Relink.Flags().StringVar(&configFlag, "conf", "", "Directory of the custom config")
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
	customPreConfig, customConfig := viper.GetString("vim.CustomPreConfig"), viper.GetString("vim.CustomConfig")
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
var Relink = &cobra.Command{
	Use:   "relink",
	Short: "Manually specify the directories you want as your preconfig and custom config.  Either are optional",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(viper.GetString("vim.CustomPreConfig"), viper.GetString("vim.CustomConfig"))
		fmt.Println(viper.GetString("vim.pre-config"), viper.GetString("vim.config"))
		return linkCustomConfig(viper.GetString("pre-config"), viper.GetString("config"))
	},
}
