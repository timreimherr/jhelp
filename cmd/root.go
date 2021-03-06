/*
Copyright © 2022 Tim Reimherr timreimherr@gmail.com

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/timreimherr/jhelp/internal/service"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jhelp",
	Short: "jhelp prints cli commands for various dev tools and allows users CRUD operations for additional tool commands",
	Long: `jhelp is a CLI library that prints cli commands
for various dev tools and allows users CRUD operations
for additional tool commands.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		section, _ := cmd.Flags().GetString("get")
		if section != "" {
			service.PrintSection(section)
		} else {
			service.PrintAll()
		}
	},
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jhelp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("get", "g", "", "Get will search for a section of it's value. It will be printed if found, otherwise all sections will be found")
}
