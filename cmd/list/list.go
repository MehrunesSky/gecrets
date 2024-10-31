/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package list

import (
	"github.com/MehrunesSky/gecrets/editors/vim"
	"github.com/MehrunesSky/gecrets/keyvaults"
	"github.com/MehrunesSky/gecrets/keyvaults/azure"
	"regexp"

	"github.com/spf13/cobra"
)

var keystore string

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Key brief description of your command",
	Long: `Key longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		v := azure.NewVault(keystore)
		r, _ := cmd.Flags().GetString("regex")
		var secretOptions *keyvaults.GetSecretsOption
		if r != "" {
			secretOptions = &keyvaults.GetSecretsOption{
				Regex: regexp.MustCompile(r),
			}
		}
		secrets, err := v.GetSecrets(secretOptions)
		if err != nil {
			return err
		}
		vim.NewVimExec(v.GetSecretModel()).Open(secrets)

		return nil
	},
}

func init() {
	ListCmd.Flags().StringP("regex", "r", "", "Regex for filter secrets")

	ListCmd.Flags().StringVar(&keystore, "ks", "", "name of keystore")
	ListCmd.MarkFlagRequired("ks")
}
