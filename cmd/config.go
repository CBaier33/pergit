/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
  "os"
  "os/user"
  "path/filepath"
  "bufio"

  "encoding/json"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure student info",
	Run: func(cmd *cobra.Command, args []string) {

    // Record student info
    var firstName string
    var lastName string
    var email string

    fmt.Print("First Name: ")
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        firstName = scanner.Text()
    }

    fmt.Print("Last Name: ")
    scanner = bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        lastName = scanner.Text()
    }

    fmt.Print("Email: ")
    scanner = bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        email = scanner.Text()
    }

    // Store student data
    student := map[string]string{
      "firstName": firstName,
      "lastName": lastName,
      "email": email,
    }

    data, err := json.Marshal(student)
    if err != nil {
      fmt.Println(err)
      return
    }

    // Get current user
   	usr, err := user.Current()
    if err != nil {
      fmt.Println(err)
      return
    }

    // save input to json file
    err = os.WriteFile(filepath.Join(usr.HomeDir, ".pergit.json"), data, 0644)
    if err != nil {
      fmt.Println(err)
      return
    }
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
