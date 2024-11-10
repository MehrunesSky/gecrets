package update

import (
	"github.com/MehrunesSky/gecrets/common"
	"github.com/MehrunesSky/gecrets/editors"
	"github.com/MehrunesSky/gecrets/keyvaults"
	"github.com/MehrunesSky/gecrets/keyvaults/azure"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUpdateCmd_shouldErrorWhenKsNotProvided(t *testing.T) {
	assert.EqualError(t, UpdateCmd.Execute(), "flag accessed but not defined: ks")
}

func TestNewUpdateCmd_shouldNoChange(t *testing.T) {

	keyVaultService := keyvaults.NewMockKeyVaultService(t)

	keyVaultService.On("GetSecrets", (*keyvaults.GetSecretsOption)(nil)).
		Return(common.SecretIs{}, nil)
	keyVaultService.On("GetSecretModel").
		Return(azure.AzureSecret{})

	editorService := editors.NewMockEditorService(t)
	editorService.On("Update", common.SecretIs{}).Return(common.SecretIs{})

	updateCmd := NewUpdateCmd(
		func(s string) keyvaults.KeyVaultService {
			return keyVaultService
		},
		func(s string, i common.SecretI) (editors.EditorService, error) {
			return editorService, nil
		},
	)

	updateCmd.Flags().String("ks", "", "")
	updateCmd.Flags().String("editor", "vim", "")

	updateCmd.SetArgs([]string{"--ks=test"})

	require.NoError(t, updateCmd.Execute())

	keyVaultService.AssertExpectations(t)
	editorService.AssertExpectations(t)

}
