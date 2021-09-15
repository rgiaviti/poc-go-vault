package main

import (
	"fmt"
	"github.com/hashicorp/vault/api"
)

func main() {
	var token = "myroot"
	var vaultHost = "http://localhost:8200"

	vaultConfig := &api.Config{
		Address: vaultHost,
	}

	vaultClient, err := api.NewClient(vaultConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	vaultClient.SetToken(token)
	secret, err := vaultClient.Logical().Read("secret/data/mytest")

	if err != nil {
		fmt.Println(err)
		return
	}

	m, ok := secret.Data["data"].(map[string]interface{})

	if !ok {
		fmt.Printf("%T %#v\n", secret.Data["data"], secret.Data["data"])
		return
	}
	fmt.Printf("hello: %v\n", m["nome"])
}

