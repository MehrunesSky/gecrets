package gsecrets

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	"log"
)

type Vault struct {
	KeyVaultName string
	client       *azsecrets.Client
}

func NewVault(keyVaultName string) Vault {
	keyVaultURL := fmt.Sprintf("https://%s.vault.azure.net/", keyVaultName)
	//Create a credential using the NewDefaultAzureCredential type.
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}

	//Establish a connection to the Key Vault client
	client, err := azsecrets.NewClient(keyVaultURL, cred, nil)
	if err != nil {
		log.Fatalf("failed to connect to client: %v", err)
	}
	return Vault{
		KeyVaultName: keyVaultName,
		client:       client,
	}
}

func (v Vault) GetSecretIds() []string {
	var ret []string
	pager := v.client.NewListSecretsPager(nil)
	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		if err != nil {
			log.Fatalln(err)
		}
		for _, secret := range resp.SecretListResult.Value {
			ret = append(ret, secret.ID.Name())
		}
	}
	return ret
}

func (v Vault) SetSecretValue(key string, value string) {
	_, err := v.client.SetSecret(context.TODO(), key, azsecrets.SetSecretParameters{
		Value: &value,
	}, nil)
	if err != nil {
		log.Fatalf("failed to create a secret: %v", err)
	}
}

func (v Vault) GetSecretValue(id string) string {
	resp, err := v.client.GetSecret(context.TODO(), id, "", nil)
	if err != nil {
		log.Fatalf("failed to get the secret: %v", err)
	}
	return *resp.Value
}
