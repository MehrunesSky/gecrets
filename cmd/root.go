package cmd

import (
	"github.com/MehrunesSky/gecrets/cmd/list"
	"github.com/MehrunesSky/gecrets/cmd/update"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gecrets",
	Short: "Key brief description of your application",
	Long: `Key longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(update.UpdateCmd)
	rootCmd.AddCommand(list.ListCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringP("ks", "", "", "name of keystore")
	err := rootCmd.MarkPersistentFlagRequired("ks")
	if err != nil {
		log.Fatalln(err)
	}

	rootCmd.PersistentFlags().StringP("editor", "e", "vim", "Editor")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
