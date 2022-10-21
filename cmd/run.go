package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Configurations for flags
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "roku remote",
		Short: "A modern Roku remote control CLI",
		Long: `Use roku remote to interact with roku devices on your network!`,
	}
)

// init CLI
func Run() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}	
}

func init() {
	cobra.OnInitialize(config)
}

func config() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Find Config
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Config:", viper.ConfigFileUsed())
	}
}