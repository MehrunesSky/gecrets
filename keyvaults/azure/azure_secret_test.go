package azure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAzureSecret_ToJson(t *testing.T) {

	secret1 := NewAzureSecret(
		"Key1",
		"Value1",
		"ContentType1",
	)

	secret2 := NewAzureSecret(
		"Key1",
		"Value2",
		"ContentType2",
	)

	assert.True(t, secret1.Diff(secret2))

}

func TestAzureSecret_GetKey(t *testing.T) {
	secret1 := NewAzureSecret(
		"Key1",
		"Value1",
		"ContentType1",
	)

	assert.Equal(t, "Key1", secret1.GetKey())
}
