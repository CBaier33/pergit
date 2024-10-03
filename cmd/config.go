/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
  "os"
  "os/user"
  "path/filepath"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "configure student information",
	Run: func(cmd *cobra.Command, args []string) {

    // Record student info
    
    var firstname string
    var lastname string
    var email string

    fmt.Print("First Name: ")
    fmt.Scan(&firstname)

    fmt.Print("Last Name: ")
    fmt.Scan(&lastname)

    fmt.Print("Email: ")
    fmt.Scan(&email)

    // Save info to config file
   	usr, err := user.Current()
    if err != nil {
      panic(err)
    }

    confFile, err := os.Create(filepath.Join(usr.HomeDir, ".pergit.conf"))
    if err != nil {
      fmt.Println(err)
      return
    }

    confFile.WriteString(fmt.Sprintf("firstname:%s\n", firstname))
    confFile.WriteString(fmt.Sprintf("lastname:%s\n", lastname))
    confFile.WriteString(fmt.Sprintf("email:%s\n", email))

    err = confFile.Close()

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
