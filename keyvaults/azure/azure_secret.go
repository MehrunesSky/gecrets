package azure

import (
	"encoding/json"
	"github.com/MehrunesSky/gecrets/common"
)

type AzureSecret struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	ContentType string `json:"contentType"`
}

func (a AzureSecret) GetKey() string {
	return a.Key
}

func NewAzureSecret(key string, value string, contentType string) AzureSecret {
	return AzureSecret{Key: key, Value: value, ContentType: contentType}
}

func (a AzureSecret) ToJson() string {
	b, _ := json.Marshal(a)
	return string(b)
}

func (a AzureSecret) Diff(otherSecret common.SecretI) bool {
	return a.Value != otherSecret.(AzureSecret).Value ||
		a.ContentType != otherSecret.(AzureSecret).ContentType
}
