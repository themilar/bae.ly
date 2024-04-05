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
	Long:  `This command accepts a single "long" link, shortens it and copies the resulting "short" link to the clipboard`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		clipboard.Init()
		link := args[0]
		exp, _ := cmd.Flags().GetInt("expiry")
		fmt.Println("shorten called", link)
		shortLink := bae.Shorten(link, exp)
		clipboard.Write(clipboard.FmtText, []byte(shortLink))
		fmt.Println(shortLink)
		// for _, v := range args {
		// }
	},
}

func init() {
	rootCmd.AddCommand(shortenCmd)
	shortenCmd.Flags().Int("expiry", 0, "set how long you want the link to be available (in hours)")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shortenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shortenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
