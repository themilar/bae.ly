/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/themilar/bae.ly/bae"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "print a list of links you've shortend",
	Long: `This command prints out the five most recent links you've shortend by default
you can modify this behaviour by providing a custom limit as an integer to the command.
for example: 'bae list 10' returns the ten most recent links you shortened`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		var limit int
		if len(args) < 1 {
			limit = 5
		} else {
			limit, _ = strconv.Atoi(args[0])
		}
		if limit < 1 {
			fmt.Print("limit must be a positive integer")
			return
		}
		fmt.Println(bae.List(limit))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
