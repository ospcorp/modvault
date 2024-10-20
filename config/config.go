package config

import (
	"encoding/json"
	"os"
	"os/user"
	"path/filepath"
	"reflect"

	"modvault/backend"
)

const defaultConfigFilename = ".modvault.config"

type ConfigChoices struct {
	DefaultBackend string `json:"backend"`
	DefaultPrefix  string `json:"prefix,omitempty"`
	DefaultTTL     string `json:"ttl"`
}

func (cc ConfigChoices) MergeOnto(target ConfigChoices) ConfigChoices {
	var mergedChoices ConfigChoices
	ccType := reflect.TypeOf(mergedChoices)
	ccVal := reflect.ValueOf(cc)
	tVal := reflect.ValueOf(target)
	resultVal := reflect.ValueOf(&mergedChoices)

	for i := 0; i < ccType.NumField(); i++ {
		var resultFieldVal string
		ccFieldVal := ccVal.Field(i)
		targetFieldVal := tVal.Field(i)

		if ccFieldVal.String() == "" {
			resultFieldVal = targetFieldVal.String()
		} else {
			resultFieldVal = ccFieldVal.String()
		}

		resultVal.Elem().Field(i).SetString(resultFieldVal)
	}

	return mergedChoices
}

type Config struct {
	Backend backend.Backend
	Prefix  string
	TTL     string
}

func getConfigFilePath() string {
	usr, _ := user.Current()
	configPath := filepath.Join(usr.HomeDir, defaultConfigFilename)
	return configPath
}

// readConfigFile takes a file path as an argument or defaults to the
// default configuration file. The file should be in JSON format.
// readConfigFile unmarshals the JSON file into a ConfigChoices struct.
// If a path is not provided and the default configuration file does
// not exist, readConfigFile calls promptGenerateConfig.
func readConfigFile() ConfigChoices {
	// TODO: take file path as argument
	f, err := os.ReadFile(getConfigFilePath())
	// TODO: if config file does not exist, prompt config creation
	if err != nil {
		panic(err)
	}

	var configFileChoices ConfigChoices

	err = json.Unmarshal(f, &configFileChoices)
	if err != nil {
		// TODO: proper error message and exit
		panic(err)
	}

	return configFileChoices
}

// GetConfig takes in config values passed in as command line arguments
// and merges them with values read from a configuration file.
func GetConfig() Config {
	// GetConfig
	// trash implementation
	configFileChoices := readConfigFile()
	backendChoice, err := backend.GetBackend(configFileChoices.DefaultBackend)
	if err != nil {
		panic(err)
	}
	return Config{Backend: backendChoice}
}
