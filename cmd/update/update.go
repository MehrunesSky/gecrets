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
		pairs, err := v.GetSecrets(secretOptions)

		if err != nil {
			return err
		}

		editor, err := editors.GetEditorByName(editorName, v.GetSecretModel())
		if err != nil {
			return err
		}
		nPairs := editor.Update(pairs)

		printSecretsChange(pairs, nPairs)

		if utils.PromptYesNo("Would you like to apply this change?") {
			for _, p := range nPairs {
				v.SetSecretValue(p)
			}
		} else {
			log.Println("NO")
		}

		return nil
	},
}

func printSecretsChange(oldSecrets, newSecrets []common.SecretI) {
	fmt.Println("You will update this secrets =>")
	oldSecretsM, newSecretsM := mapSecretsToMap(oldSecrets), mapSecretsToMap(newSecrets)

	for k, secret := range newSecretsM {
		o, ok := oldSecretsM[k]
		if !ok {
			fmt.Println("Secrets", secret.GetKey(), "is new with this value :", secret)
		} else if o.Diff(secret) {
			fmt.Println("Secrets", secret.GetKey(), "change from", o.ToJson(), "to", secret.ToJson())
		}

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
