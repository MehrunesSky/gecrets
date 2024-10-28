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
	GetSecrets(getSecretOption *GetSecretsOption) ([]common.Secret, error)
	SetSecretValue(key string, value string) error
	GetSecretValue(id string) (string, error)
}
