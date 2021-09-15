package main

import (
	"fmt"
	"poc-go-vault/secret"
)

func main() {
	//testReadSecret()
	testReadSecrets()
	//testReadAllSecrets()
}

func testReadAllSecrets() {
	vault := &secret.Vault{
		Host:  "http://localhost:8200",
		Token: "myroot",
	}

	err := vault.New()
	if err != nil {
		return
	}

	secrets, err := vault.ReadAllSecrets("secret/data/mytest")
	if err != nil {
		return
	}

	fmt.Println(secrets)
}

func testReadSecrets() {
	vault := &secret.Vault{
		Host:  "http://localhost:8200",
		Token: "myroot",
	}

	err := vault.New()
	if err != nil {
		return
	}

	keys := []string{"nome", "b"}
	secrets, err := vault.ReadSecrets(keys, "secret/data/mytest")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(secrets)
}

func testReadSecret() {
	vault := &secret.Vault{
		Host:  "http://localhost:8200",
		Token: "myroot",
	}

	err := vault.New()
	if err != nil {
		return
	}

	s, err := vault.ReadSecret("nome", "secret/data/mytest")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(s)
}
