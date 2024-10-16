package backend

import "fmt"

// Backend should be an interface that all backends implement
type Backend int

const (
	Bitwarden Backend = iota
	Vault
)

var backendNames = map[Backend]string{
	Bitwarden: "Bitwarden",
	Vault:     "Hashicorp Vault",
}

func (b Backend) String() string {
	return backendNames[b]
}

func GetBackend(backendName string) (Backend, error) {
	for k, v := range backendNames {
		if backendName == v {
			return k, nil
		}
	}
	return -1, fmt.Errorf("No such backend as: %s", backendName)
}
