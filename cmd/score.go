/*
Copyright © 2024 Christopher Baier
*/
package cmd

import (
	"fmt"
	"os"
  "io"
  "regexp"

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
      return
    }

    file, err := os.Open(args[0])
    if err != nil{
      fmt.Println(err)
      return
    }

    defer file.Close()

    content, err := io.ReadAll(file)
    if err != nil {
      fmt.Println("Error Reading File: ", err)
      return
    }

    code := string(content)

    headerCheck := false
    typeCheck := false

    functionsRe := regexp.MustCompile(`def\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*\(`)
    functions := functionsRe.FindAllStringSubmatch(code, -1)

    if (len(functions) < 2) {
      fmt.Println("Your file contains less than two functions, therefore a header is not neccesary.")
    } else {
      headerRe := regexp.MustCompile(`#\s*`)
      headerBlocks := headerRe.FindAllString(code, -1)
      if (len(headerBlocks)) < 2 {
        fmt.Println("Your source file is missing a header.")
      } else {
        fmt.Println("Your file contains a header block.")
        headerCheck = true
      }
    }

    sigsRe := regexp.MustCompile(`def\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*\(.*?\)\s*:\s*([^->]|->\s*$)`)
    missingFuncs := sigsRe.FindAllStringSubmatch(code, -1)

    if len(missingFuncs) < 1 {
      fmt.Println("All functions match the programming style guide.")
      typeCheck = true
    } else {
      fmt.Println("The following functions are missing type signatures:")
      for _, function := range missingFuncs {
        fmt.Println("  -", function[1])
      }
    }

    if (headerCheck && typeCheck) {
      fmt.Println("\nYour source file adheres to the progamming style guide!")
    }
	},
}

func init() {
	rootCmd.AddCommand(scoreCmd)
}
