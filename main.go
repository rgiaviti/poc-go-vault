package main

import (
	"fmt"
	"github.com/rgiaviti/poc-go-vault/secret"
	"github.com/rgiaviti/poc-go-vault/serialization"
	"log"
)

/*-----------------------------------------------------------------------------
  Vault Secrets Creation

curl --location --request POST 'http://localhost:8200/v1/secret/data/person' \
--header 'X-Vault-Token: myroot' \
--header 'Content-Type: application/json' \
--data-raw '{
  "options": {
    "cas": 0
  },
  "data": {
    "LASTNAME": "Doe",
    "HOBBY_NAME": "Programming",
    "FACEBOOK_URL": "facebook/johndoe"
  }
}'
--------------------------------------------------------------------------------*/

func main() {

	// SetUp Vault Client
	vault := &secret.Vault{
		Host:  "http://localhost:8200",
		Token: "myroot",
	}

	// Open Vault
	err := vault.OpenVault()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Read all secrets in this specific path
	secretsMap, err := vault.ReadAllSecretsAsStringMap("secret/data/person")
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}

	// Now load the yaml as string
	data, err := serialization.ReadExampleFile()
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}

	fmt.Println(secret.HasToken(data))

	// Let's replace all secrets in loaded yaml
	replacedData, err := secret.ReplaceTokens(secretsMap, data)
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}

	// Now serialize the new replaced data to Person
	person, err := serialization.NewPerson(replacedData)
	if err != nil {
		log.Fatalf("error: %v", err)
		return
	}

	println(person.Name)
	println(person.SocialNetworks[1].Url)
}
