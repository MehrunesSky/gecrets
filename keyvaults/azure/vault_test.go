package azure

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
	"github.com/MehrunesSky/gecrets/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestVault_GetSecretIds(t *testing.T) {

	client := newMockClient(t)

	v := Vault{
		KeyVaultName: "test",
		client:       client,
	}

	client.On("NewListSecretsPager", (*azsecrets.ListSecretsOptions)(nil)).
		Return(runtime.NewPager(runtime.PagingHandler[azsecrets.ListSecretsResponse]{
			More: func(response azsecrets.ListSecretsResponse) bool {
				return false
			},
			Fetcher: func(ctx context.Context, a *azsecrets.ListSecretsResponse) (azsecrets.ListSecretsResponse, error) {
				contentType := "k8s-secret"
				id := "https://myvaultname.vault.azure.net/keys/key1053998307/b86c2e6ad9054f4abf69cc185b99aa60"
				return azsecrets.ListSecretsResponse{
					SecretListResult: azsecrets.SecretListResult{
						NextLink: nil,
						Value: []*azsecrets.SecretItem{
							{
								Attributes:  nil,
								ContentType: &contentType,
								ID:          (*azsecrets.ID)(&id),
								Tags:        nil,
								Managed:     nil,
							},
						},
					},
				}, nil
			},
			Tracer: tracing.Tracer{},
		}))
	_ = v

	secrets := v.GetSecretIds()

	assert.Equal(t, []string{"key1053998307"}, secrets)

	client.AssertExpectations(t)

}

func TestVault_getSecretsFromPage(t *testing.T) {
	client := newMockClient(t)

	value := "NewValue"
	client.On("GetSecret", context.TODO(), "key1053998307", "", (*azsecrets.GetSecretOptions)(nil)).
		Return(azsecrets.GetSecretResponse{
			SecretBundle: azsecrets.SecretBundle{
				Attributes:  nil,
				ContentType: nil,
				ID:          nil,
				Tags:        nil,
				Value:       &value,
				Kid:         nil,
				Managed:     nil,
			},
		}, nil)

	v := Vault{
		KeyVaultName: "test",
		client:       client,
	}

	contentType := "k8s-secret"
	id := "https://myvaultname.vault.azure.net/keys/key1053998307/b86c2e6ad9054f4abf69cc185b99aa60"
	secrets := azsecrets.ListSecretsResponse{
		SecretListResult: azsecrets.SecretListResult{
			NextLink: nil,
			Value: []*azsecrets.SecretItem{
				{
					Attributes:  nil,
					ContentType: &contentType,
					ID:          (*azsecrets.ID)(&id),
					Tags:        nil,
					Managed:     nil,
				},
			},
		},
	}

	secretsFromPage, err := v.getSecretsFromPage(secrets, nil)

	require.NoError(t, err)
	assert.Equal(t, common.SecretIs{
		AzureSecret{Key: "key1053998307", Value: "NewValue", ContentType: "k8s-secret"}}, secretsFromPage)
}
