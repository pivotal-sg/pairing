package config

import (
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/pivotal-sg/pairing/config"
	"github.com/spf13/viper"
)

var (
	customPreConfig = "$HOME/.config/nvim/custom_preconfig"
	customConfig    = "$HOME/.config/nvim/custom_config"
)

func init() {
	var err error
	customPreConfig, err = filepath.Abs(os.ExpandEnv(customPreConfig))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%g\n", err)
		os.Exit(1)
	}
	customConfig, err = filepath.Abs(os.ExpandEnv(customConfig))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%g\n", err)
		os.Exit(1)
	}

	viper.SetDefault("vim.CustomPreConfig", customPreConfig)
	viper.SetDefault("vim.CustomConfig", customConfig)
}
