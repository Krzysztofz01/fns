package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path"
	"sync"

	"github.com/Krzysztofz01/fns/utils"
)

const (
	configFileName string = ".fns"
	configFileType string = "json"
)

var (
	instance     *Configuration
	instanceOnce sync.Once
)

func GetConfiguration() *Configuration {
	instanceOnce.Do(func() {
		instance = loadConfiguration()
	})

	return instance
}

type Configuration struct {
	NoteReadDirectoryPaths []string `mapstructure:"note-read-directory-paths"`
	NoteWriteDirectoryPath string   `mapstructure:"note-write-directory-path"`
	EditorPath             string   `mapstructure:"editor-path"`
	TrimNote               bool     `mapstructure:"trim-note"`
}

func loadConfiguration() *Configuration {
	configPath, configExists := getLocalConfigFilePath()
	if !configExists {
		configPath, configExists = getHomeConfigFilePath()
	}

	viper.SetConfigFile(configPath)
	viper.SetConfigType(configFileType)

	viper.SetDefault("note-read-directory-paths", []string{})
	viper.SetDefault("note-write-directory-path", "")
	viper.SetDefault("editor-path", "")
	viper.SetDefault("trim-note", true)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("config: failed to read the config file: %w", err))
	}

	configuration := new(Configuration)
	if err := viper.Unmarshal(&configuration); err != nil {
		panic(fmt.Errorf("config: failed to unmarshal the config file content: %w", err))
	}

	return configuration
}

func getLocalConfigFilePath() (string, bool) {
	cwd, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("config: failed to access the current working directory: %w", err))
	}

	p := path.Join(cwd, configFileName)
	ok, err := utils.FileExist(p)
	if err != nil {
		panic(fmt.Errorf("config: failed to determine if the local config file exists: %w", err))
	}

	return p, ok
}

func getHomeConfigFilePath() (string, bool) {
	cwd, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("config: failed to access the users home directory path: %w", err))
	}

	p := path.Join(cwd, configFileName)
	ok, err := utils.FileExist(p)
	if err != nil {
		panic(fmt.Errorf("config: failed to determine if the home config file exists: %w", err))
	}

	return p, ok
}
