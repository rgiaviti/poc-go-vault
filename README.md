# POC-Go-Vault
This is a simple POC reading Secrets in Vault Servers using the oficial client.

## References:
- Vault Server: https://www.vaultproject.io/
- Client: https://pkg.go.dev/github.com/hashicorp/vault/api#section-readme

## Inserting Secret Data in Vault
```bash
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

```
