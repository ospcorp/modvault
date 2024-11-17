package main

import (
	"fmt"

	"modvault/backend"
	"modvault/config"
)

/* Proposed project structure:

|- modvault
    |- backend
        |- backend.go
        |- bitwarden.go
        |- vault.go
    |- config
        |- config.go
    |- flags (using Cobra)
    |- keys
    |- token (if necessary)
    |- modvault.go
    |- go.mod
    |- go.sum
    |- .gitignore
    |- README.md

*/

// main gets flags, determines what backend is used, passes flags to backend module via function call
func main() {
	configuration := config.GetConfig(backend.Configuration{})
	fmt.Printf("Backend: %s", configuration.BackendName)
}
