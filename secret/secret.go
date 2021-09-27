package secret

import (
	"errors"
	"fmt"
	"github.com/hashicorp/vault/api"
)

const (
	JSONDataAttr = "data"
)

type Vault struct {
	Host   string
	Token  string
	client *api.Client
}

func (v *Vault) OpenVault() error {
	apiConfig := api.Config{Address: v.Host}
	client, err := api.NewClient(&apiConfig)
	if err != nil {
		return err
	}
	v.client = client
	v.client.SetToken(v.Token)
	return nil
}

func (v *Vault) ReadSecret(key string, secretPath string) (string, error) {
	vaultSecrets, err := v.ReadAllSecrets(secretPath)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v", vaultSecrets[key]), nil
}

func (v *Vault) ReadSecrets(keys []string, secretPath string) (map[string]string, error) {
	vaultSecrets, err := v.ReadAllSecrets(secretPath)
	if err != nil {
		return nil, err
	}

	secretMap := make(map[string]string)
	for _, key := range keys {
		if value, ok := vaultSecrets[key]; ok {
			secretMap[key] = fmt.Sprintf("%v", value)
		}
	}

	return secretMap, nil
}

func (v *Vault) ReadAllSecretsAsStringMap(secretPath string) (map[string]string, error) {
	vaultSecrets, err := v.ReadAllSecrets(secretPath)
	if err != nil {
		return nil, err
	}

	secretMap := make(map[string]string)

	for key, value := range vaultSecrets {
		secretMap[key] = fmt.Sprintf("%v", value)
	}

	return secretMap, nil
}

func (v *Vault) ReadAllSecrets(secretPath string) (map[string]interface{}, error) {
	if v.client == nil {
		return nil, errors.New("vault client not initialized. you should call OpenVault() first")
	}

	vaultSecrets, err := v.client.Logical().Read(secretPath)
	if err != nil {
		return nil, err
	}

	vaultData, ok := vaultSecrets.Data[JSONDataAttr].(map[string]interface{})
	if !ok {
		return nil, err
	}
	return vaultData, nil
}
