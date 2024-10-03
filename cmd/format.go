/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// formatCmd represents the format command
var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Format file according to the Programming Style Guide",
	Run: func(cmd *cobra.Command, args []string) {

    // initial checks
    if (len(args) != 1) {
      fmt.Println("Incorrect arguments. Please enter a python file to be formatted.")
    } else if !(strings.HasSuffix(args[0], ".py")) {
      fmt.Println("Please enter a python file to be formatted.")
    } else {

      // check for config file

      // get file description

      // get class

      // assign date / due date
      
      // pull student info from conf file

      // insert code block into file
      fmt.Println("Formatted.")
    }
	},
}

func init() {
	rootCmd.AddCommand(formatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// formatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// formatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
