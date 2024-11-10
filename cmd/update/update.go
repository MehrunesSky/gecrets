/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package update

import (
	"fmt"
	"github.com/MehrunesSky/gecrets/common"
	"github.com/MehrunesSky/gecrets/editors"
	"github.com/MehrunesSky/gecrets/keyvaults"
	"github.com/MehrunesSky/gecrets/keyvaults/azure"
	"github.com/MehrunesSky/gecrets/utils"
	"github.com/spf13/cobra"
	"log"
	"regexp"
)

var regex string

// UpdateCmd represents the update command
var UpdateCmd = NewUpdateCmd(
	func(s string) keyvaults.KeyVaultService {
		return azure.NewVault(s)
	},
	func(s string, model common.SecretI) (editors.EditorService, error) {
		return editors.GetEditorByName(s, model)
	},
)

func NewUpdateCmd(keyVaultService func(string) keyvaults.KeyVaultService,
	editorService func(string, common.SecretI) (editors.EditorService, error)) *cobra.Command {

	return &cobra.Command{
		Use:   "update",
		Short: "Key brief description of your command",
		Long: `Key longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ks, err := cmd.Flags().GetString("ks")
			if err != nil {
				return err
			}

			editorName, err := cmd.Flags().GetString("editor")
			if err != nil {
				return err
			}
			v := keyVaultService(ks)
			var secretOptions *keyvaults.GetSecretsOption
			if regex != "" {
				secretOptions = &keyvaults.GetSecretsOption{Regex: regexp.MustCompile(regex)}
			}
			oldSecrets, err := v.GetSecrets(secretOptions)

			if err != nil {
				return err
			}

			editor, err := editorService(editorName, v.GetSecretModel())
			if err != nil {
				return err
			}
			newSecrets := editor.Update(oldSecrets)

			new, changed := oldSecrets.GetChangedSecrets(newSecrets)

			if len(new) == 0 && len(changed) == 0 {
				fmt.Println("No change")
				return nil
			}

			printSecretsChange(new, changed)

			if utils.PromptYesNo("Would you like to apply this change?") {
				for _, p := range newSecrets {
					v.SetSecretValue(p)
				}
			} else {
				log.Println("NO")
			}

			return nil
		},
	}
}

func printSecretsChange(newSecrets, updateSecrets []common.Changed) {
	fmt.Println("You will update this secrets =>")

	for _, secret := range newSecrets {
		fmt.Println("Secrets", secret.Key, "is new with this value :", secret.NewValue)
	}

	for _, secret := range updateSecrets {
		fmt.Println("Secrets", secret.Key, "change from", secret.OldValue, "to", secret.NewValue)
	}

}

func init() {

	UpdateCmd.Flags().StringVarP(&regex, "regex", "r", "", "Help message for toggle")
}
