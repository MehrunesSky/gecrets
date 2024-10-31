package keyvaults

import (
	"github.com/MehrunesSky/gecrets/common"
	"regexp"
)

type GetSecretsOption struct {
	Regex *regexp.Regexp
}

type KeyVaultService interface {
	GetSecretIds() []string
	GetSecrets(getSecretOption *GetSecretsOption) ([]common.SecretI, error)
	SetSecretValue(secretI common.SecretI) error
	GetSecretValue(id string) (string, error)
	GetSecretModel() common.SecretI
}
