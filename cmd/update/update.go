/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package update

import (
	"fmt"
	"github.com/MehrunesSky/gecrets/common"
	"github.com/MehrunesSky/gecrets/editors/vim"
	"github.com/MehrunesSky/gecrets/keyvaults"
	"github.com/MehrunesSky/gecrets/keyvaults/azure"
	"github.com/MehrunesSky/gecrets/utils"
	"github.com/spf13/cobra"
	"log"
	"regexp"
)

var keystore string

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
		v := azure.NewVault("mehr")
		var secretOptions *keyvaults.GetSecretsOption
		if regex != "" {
			secretOptions = &keyvaults.GetSecretsOption{Regex: regexp.MustCompile(regex)}
		}
		pairs, err := v.GetSecrets(secretOptions)

		if err != nil {
			return err
		}

		nPairs := vim.NewVimExec().UpdateWithVim(pairs)

		printSecretsChange(pairs, nPairs)

		if utils.PromptYesNo("Would you like to apply this change?") {
			for _, p := range nPairs {
				v.SetSecretValue(p.Key, p.Value)
			}
		} else {
			log.Println("NO")
		}

		return nil
	},
}

func printSecretsChange(oldSecrets, newSecrets []common.Secret) {
	fmt.Println("You will update this secrets =>")
	oldSecretsM, newSecretsM := mapSecretsToMap(oldSecrets), mapSecretsToMap(newSecrets)

	for k, secret := range newSecretsM {

		o, ok := oldSecretsM[k]
		if !ok {
			fmt.Println("Secrets", secret.Key, "is new with this value :", secret.Value)
		} else if o.Value != secret.Value {
			fmt.Println("Secrets", secret.Key, "change from", o.Value, "to", secret.Value)
		}

	}

}

func mapSecretsToMap(secrets []common.Secret) map[string]common.Secret {

	m := make(map[string]common.Secret)

	for _, secret := range secrets {
		m[secret.Key] = secret
	}

	return m

}

func init() {

	UpdateCmd.Flags().StringVarP(&regex, "regex", "r", "", "Help message for toggle")

	UpdateCmd.Flags().StringVar(&keystore, "ks", "", "name of keystore")
	UpdateCmd.MarkFlagRequired("ks")
}
