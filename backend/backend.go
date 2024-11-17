package backend

import (
	"fmt"
	"reflect"
)

// the return values might turn into some kind of Secret struct
type Backend interface {
	AddSecret(string) error
	ListSecrets() ([]string, error)
	SearchSecret(string) ([]string, error)
	GetSecret(string) string
}

var backends = map[string]Backend{
	// something like this?
	"vault": VaultBackend{},
	// "bitwarden": BitwardenBackend,
	// "s3": S3Backend,
}

var backendNames = []string{}

func ListBackendNames() []string {
	if len(backendNames) == 0 {
		initBackendNames()
	}
	return backendNames
}

func initBackendNames() {
	for k := range backends {
		backendNames = append(backendNames, k)
	}
}

func GetBackend(configuration Configuration) (Backend, error) {
	switch backendName := configuration.BackendName; backendName {
	case "vault":
		return VaultBackend{configuration}, nil
	default:
		return nil, fmt.Errorf("No backend found for: %s", configuration.BackendName)
	}
}

type Configuration struct {
	BackendName string `json:"backend"`
	Prefix      string `json:"prefix,omitempty"`
	TTL         string `json:"ttl"`
}

func (cc Configuration) MergeOnto(target Configuration) Configuration {
	var mergedChoices Configuration
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
