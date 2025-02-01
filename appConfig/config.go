package appConfig

import (
	"github.com/spf13/viper"
)

type Config struct {
	ObsidianPath string `mapstructure:"ObsidianPath"`
	GithubRepo   string `mapstructure:"GithubRepo"`
}

func InitConfig() {
	viper.SetDefault("ObsidianPath", "/home/c/obsidian-test/")
	viper.SetDefault("GithubRepo", "None")

	viper.AddConfigPath("~/.config/tidymd")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		println("Pas de fichier de config, utilise les defaults")
	}
}
