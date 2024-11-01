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
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Key brief description of your command",
	Long: `Key longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ks, _ := cmd.Flags().GetString("ks")
		editorName, _ := cmd.Flags().GetString("editor")
		v := azure.NewVault(ks)
		var secretOptions *keyvaults.GetSecretsOption
		if regex != "" {
			secretOptions = &keyvaults.GetSecretsOption{Regex: regexp.MustCompile(regex)}
		}
		oldSecrets, err := v.GetSecrets(secretOptions)

		if err != nil {
			return err
		}

		editor, err := editors.GetEditorByName(editorName, v.GetSecretModel())
		if err != nil {
			return err
		}
		newSecrets := editor.Update(oldSecrets)

		new, changed := getChangedSecret(oldSecrets, newSecrets)

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

type Changed struct {
	Key      string
	OldValue string
	NewValue string
	New      bool
}

func NewChanged(key string, oldValue string, newValue string, new bool) Changed {
	return Changed{Key: key, OldValue: oldValue, NewValue: newValue, New: new}
}

func getChangedSecret(oldSecrets, newSecrets []common.SecretI) (new []Changed, changed []Changed) {
	oldSecretsM, newSecretsM := mapSecretsToMap(oldSecrets), mapSecretsToMap(newSecrets)

	for k, secret := range newSecretsM {
		o, ok := oldSecretsM[k]
		if !ok {
			fmt.Println("Secrets", secret.GetKey(), "is new with this value :", secret)
			new = append(
				new,
				NewChanged(secret.GetKey(), "", secret.ToJson(), true),
			)
		} else if o.Diff(secret) {
			changed = append(
				changed,
				NewChanged(secret.GetKey(), o.ToJson(), secret.ToJson(), false),
			)
		}

	}
	return new, changed
}

func printSecretsChange(newSecrets, updateSecrets []Changed) {
	fmt.Println("You will update this secrets =>")

	for _, secret := range newSecrets {
		fmt.Println("Secrets", secret.Key, "is new with this value :", secret.NewValue)
	}

	for _, secret := range updateSecrets {
		fmt.Println("Secrets", secret.Key, "change from", secret.OldValue, "to", secret.NewValue)
	}

}

func mapSecretsToMap(secrets []common.SecretI) map[string]common.SecretI {

	m := make(map[string]common.SecretI)

	for _, secret := range secrets {
		m[secret.GetKey()] = secret
	}

	return m
}

func init() {
	UpdateCmd.Flags().StringVarP(&regex, "regex", "r", "", "Help message for toggle")
}
