package keyvaults

import (
	"github.com/MehrunesSky/gecrets/common"
	"regexp"
)

type GetSecretsOption struct {
	Regex *regexp.Regexp
}

//go:generate go run -mod=mod github.com/vektra/mockery/v2 --name KeyVaultService --inpackage --inpackage-suffix
type KeyVaultService interface {
	GetSecretIds() []string
	GetSecrets(getSecretOption *GetSecretsOption) (common.SecretIs, error)
	SetSecretValue(secretI common.SecretI) error
	GetSecretValue(id string) (string, error)
	GetSecretModel() common.SecretI
}
