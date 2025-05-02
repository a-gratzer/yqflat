package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

func InitDefaultViperConfig(cfgFile string) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, _ := os.UserHomeDir()

		// Search config in home directory with name ".searchHero" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".flatternYaml")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	//// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Using config file failed:", viper.ConfigFileUsed())
	}
}

func ReadFile(path string) string {

	content, error := os.ReadFile(path)

	if error != nil {
		log.Fatal(error)
	}

	return string(content)
}
