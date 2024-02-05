/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listTopCmd represents the listTop command
var listTopCmd = &cobra.Command{
	Use:   "listTop",
	Short: "It returns a list of the most popular songs",
	Long: `
	You can use this command to get a list of the most popular songs in the world.
	You can also use the --country flag to get the most popular songs in a specific country.
	For example:
		$ spotify listTop --country=FR
		$ spotify listTop --country=ES

	You can also use the --limit flag to limit the number of songs returned. By default, it returns 20 songs.
	For example:
		$ spotify listTop --limit=10
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listTop called")
	},
}

func init() {
	rootCmd.AddCommand(listTopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listTopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listTopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
