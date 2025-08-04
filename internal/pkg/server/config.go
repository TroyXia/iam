package server

import (
	"github.com/TroyXia/iam/pkg/log"
	"github.com/marmotedu/component-base/pkg/util/homedir"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

const (
	// RecommendedHomeDir defines the default directory used to place all iam service configurations.
	RecommendedHomeDir = ".iam"

	// RecommendedEnvPrefix defines the ENV prefix used by all iam service.
	RecommendedEnvPrefix = "IAM"
)

// LoadConfig reads in config file and ENV variables if set.
func LoadConfig(cfg string, defaultName string) {
	if cfg != "" {
		viper.SetConfigFile(cfg)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath(filepath.Join(homedir.HomeDir(), RecommendedHomeDir))
		viper.AddConfigPath("/etc/iam")
		viper.SetConfigName(defaultName)
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix(RecommendedEnvPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Warnf("WARNING: viper failed to discover and load the configuration file: %s", err.Error())
	}
}
