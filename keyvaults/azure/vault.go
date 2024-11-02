package azure

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	"github.com/MehrunesSky/gecrets/common"
	"github.com/MehrunesSky/gecrets/keyvaults"
	"log"
)

type Vault struct {
	KeyVaultName string
	client       *azsecrets.Client
}

func (v Vault) GetSecretModel() common.SecretI {
	return &AzureSecret{}
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

func (v Vault) GetSecrets(getSecretOption *keyvaults.GetSecretsOption) (common.SecretIs, error) {
	var ret []common.SecretI
	pager := v.client.NewListSecretsPager(nil)
	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		if err != nil {
			return nil, err
		}
		for _, secret := range resp.SecretListResult.Value {
			secretName := secret.ID.Name()
			if getSecretOption == nil || getSecretOption.Regex.MatchString(secretName) {
				value, err := v.GetSecretValue(secretName)
				if err != nil {
					return nil, err
				}
				contentType := ""

				if secret.ContentType != nil {
					contentType = *secret.ContentType
				}

				ret = append(
					ret,
					NewAzureSecret(secretName, value, contentType),
				)
			}
		}
	}
	return ret, nil
}

func (v Vault) SetSecretValue(secretI common.SecretI) error {
	secret := secretI.(AzureSecret)
	_, err := v.client.SetSecret(
		context.TODO(),
		secret.Key,
		azsecrets.SetSecretParameters{
			Value:       &secret.Value,
			ContentType: &secret.ContentType,
		}, nil)
	return err
}

func (v Vault) GetSecretValue(id string) (string, error) {
	resp, err := v.client.GetSecret(context.TODO(), id, "", nil)
	return *resp.Value, err
}
