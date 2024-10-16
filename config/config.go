package config

import (
	"encoding/json"
	"os"
	"os/user"
	"path/filepath"

	"modvault/backend"
)

const defaultConfigFilename = ".modvault.config"

type configJson struct {
	BackendName string `json:"backend"`
}

type Config struct {
	Backend backend.Backend
}

func getConfigFilePath() string {
	usr, _ := user.Current()
	configPath := filepath.Join(usr.HomeDir, defaultConfigFilename)
	return configPath
}

// readConfigFile takes a file path as an argument or defaults to the
// default configuration file. The file should be in JSON format.
// readConfigFile unmarshals the JSON file into a configJson struct.
// If a path is not provided and the default configuration file does
// not exist, readConfigFile calls promptGenerateConfig.
func readConfigFile() configJson {
	// TODO: take file path as argument
	f, err := os.ReadFile(getConfigFilePath())
	// TODO: if config file does not exist, prompt config creation
	if err != nil {
		panic(err)
	}

	var configFileJson configJson

	err = json.Unmarshal(f, &configFileJson)
	if err != nil {
		panic(err)
	}

	return configFileJson
}

// GetConfig takes in config values passed in as command line arguments
// and merges them with values read from a configuration file.
func GetConfig() Config {
	// GetConfig
	// trash implementation
	configFileJson := readConfigFile()
	backendChoice, err := backend.GetBackend(configFileJson.BackendName)
	if err != nil {
		panic(err)
	}
	return Config{Backend: backendChoice}
}
