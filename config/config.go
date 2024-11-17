package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"modvault/backend"
)

const (
	defaultConfigFilename = ".modvault.config"
)

var inputReader = bufio.NewReader(os.Stdin)

func getConfigFilePath() string {
	usr, err := user.Current()
	if err != nil {
		// TODO: better error handling
		panic(err)
	}
	configPath := filepath.Join(usr.HomeDir, defaultConfigFilename)
	return configPath
}

// readConfigFile takes a file path as an argument or defaults to the
// default configuration file. The file should be in JSON format.
// readConfigFile unmarshals the JSON file into a Configuration struct.
// If a path is not provided and the default configuration file does
// not exist, readConfigFile calls promptGenerateConfig.
func readConfigFile() backend.Configuration {
	// TODO: take file path as argument
	f, err := os.ReadFile(getConfigFilePath())
	// TODO: if config file does not exist, prompt config creation
	if err != nil {
		panic(err)
	}

	var configFileChoices backend.Configuration

	err = json.Unmarshal(f, &configFileChoices)
	if err != nil {
		// TODO: proper error message and exit
		panic(err)
	}

	return configFileChoices
}

func promptForConfiguration() backend.Configuration {
	var choices backend.Configuration
	// TODO: refactor this into something nicer if it makes sense
	backendNames := strings.Join(backend.ListBackendNames(), ", ")
	backendPrompt := fmt.Printf("Please define the default backend (%s):\n> ", backendNames)
	choices.BackendName = promptForInput(backendPrompt, backendValidator)

	return choices
}

type validator func(string) bool

func promptForInput(prompt string, fn validator) string {
	fmt.Println(prompt)
	text, _ := inputReader.ReadString('\n')
	if fn(text) {
		return text
	}
	return ""
}

func backendValidator(backendName string) bool {
	return false
}

// GetConfig takes in config values passed in as command line arguments
// and merges them with values read from a configuration file.
func GetConfig(argumentChoices backend.Configuration) backend.Configuration {
	configFileChoices := readConfigFile()
	mergedChoices := argumentChoices.MergeOnto(configFileChoices)
	return mergedChoices
}
