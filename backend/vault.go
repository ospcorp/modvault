package backend

type VaultBackend struct {
	configuration Configuration
}

func (VaultBackend) AddSecret(key string) error {
	return nil
}

func (VaultBackend) ListSecrets() ([]string, error) {
	return nil, nil
}

func (VaultBackend) SearchSecret(key string) ([]string, error) {
	return nil, nil
}

func (VaultBackend) GetSecret(key string) string {
	return ""
}
