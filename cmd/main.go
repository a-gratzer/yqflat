package main

import (
	"flag"
	"github.com/a-gratzer/qyflat/internal/config"
	internal_logger "github.com/a-gratzer/qyflat/internal/logger"
	"github.com/a-gratzer/qyflat/internal/yqflat"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var logger *zap.Logger = nil

/*
init initializes the application by parsing the configuration file flag,
setting up the default Viper configuration, and initializing the global logger.
*/
func init() {

	var configFile = flag.String("config", "./config/config.yaml", "Set path to the configuration file")
	flag.Parse()

	config.InitDefaultViperConfig(*configFile)
	logger = internal_logger.GetZapLogger(false, false)
}

/*
main is the entry point of the application. It reads the YAML file path from command-line flags,
retrieves the configured separator, parses and flattens the YAML file using the internal yqflat package,
and logs each flattened key-value pair using the global logger.
*/
func main() {

	var filePath = mustReadFilePath()
	var flatYaml map[string]string
	var err error
	var separator = viper.GetViper().GetString(config.CONFIG_SEPARATOR)

	if flatYaml, err = yqflat.New(separator).Parse(filePath); err != nil {
		panic(err)
	}

	for k, v := range flatYaml {
		logger.Debug("yqflat", zap.Any(k, v))
	}
}

func mustReadFilePath() string {
	var yamlFilePath = flag.String("yaml", "./example/example.yaml", "Set path to the yaml file")
	flag.Parse()
	return *yamlFilePath
}
