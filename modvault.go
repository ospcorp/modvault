package main

import (
	"fmt"

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
    |- token (if necessary)
    |- modvault.go
    |- go.mod
    |- go.sum
    |- .gitignore
    |- README.md

*/

// main gets flags, determines what backend is used, passes flags to backend module via function call
func main() {
	config := config.GetConfig()
	fmt.Printf("Backend: %s", config.Backend)
}
