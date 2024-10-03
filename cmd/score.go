/*
Copyright © 2024 Christopher Baier

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// scoreCmd represents the score command
var scoreCmd = &cobra.Command{
	Use:   "score",
	Short: "Recommended changes to match programming style guide",
	Run: func(cmd *cobra.Command, args []string) {

    // initial checks
    if (len(args) < 1) {
      fmt.Println("Please enter a python file to be scored.")
    }
    // Check for type annotations

    // Check for presence of header (needed if there is more than one function def)
	},
}

func init() {
	rootCmd.AddCommand(scoreCmd)
}
