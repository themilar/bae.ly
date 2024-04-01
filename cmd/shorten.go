/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/themilar/bae.ly/bae"
	"golang.design/x/clipboard"
)

// shortenCmd represents the shorten command
var shortenCmd = &cobra.Command{
	Use:   "shorten",
	Short: "shorten links",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		clipboard.Init()
		// bae.Shorten()
		for _, v := range args {
			fmt.Println("shorten called", v)
			shortLink := bae.Shorten(v)
			clipboard.Write(clipboard.FmtText, []byte(shortLink))
			fmt.Println(shortLink)
		}
	},
}

func init() {
	rootCmd.AddCommand(shortenCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shortenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shortenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
