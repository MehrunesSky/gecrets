package azure

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

//go:generate go run -mod=mod github.com/vektra/mockery/v2 --name client --inpackage --inpackage-suffix --testonly
type client interface {
	NewListSecretsPager(options *azsecrets.ListSecretsOptions) *runtime.Pager[azsecrets.ListSecretsResponse]
	SetSecret(todo context.Context, key string,
		parameters azsecrets.SetSecretParameters, options *azsecrets.SetSecretOptions) (azsecrets.SetSecretResponse, error)
	GetSecret(ctx context.Context, name string, version string, options *azsecrets.GetSecretOptions) (azsecrets.GetSecretResponse, error)
}

type azureClient struct {
	*azsecrets.Client
}
