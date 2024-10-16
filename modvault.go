package main

import (
	"fmt"

	"modvault/config"
)

// configuration
// tokens
//  - handles token caching
// vault
// bitwarden

// main gets flags, determines what backend is used, passes flags to backend module via function call
func main() {
	config := config.GetConfig()
	fmt.Printf("Backend: %s", config.Backend)
}
