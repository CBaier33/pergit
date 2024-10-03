/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"
  "os"
  "os/user"
  "path/filepath"
  "encoding/json"
  "bufio"

	"github.com/spf13/cobra"
)

// addheaderCmd represents the format command
var addheaderCmd = &cobra.Command{
	Use:   "addheader",
	Short: "Insert source file header.",
	Run: func(cmd *cobra.Command, args []string) {

    // initial checks
    if (len(args) != 1 || !(strings.HasSuffix(args[0], ".py")) ) {
      fmt.Println("Please enter a python file to be formatted")
      return
    }

    usr, err := user.Current()
    if err != nil {
      fmt.Println(err)
      return
    }
    
    // check for config file
    confFile, err := os.ReadFile(filepath.Join(usr.HomeDir, ".pergit.json"))

    if err != nil {
      fmt.Println("Please create a config file before formatting -> \"pergit config\".")
      return
    }

    var (
      assignName string
      description string
      class string
      professor string
      dueDate string
      assignDate string
    )

    // get file description
    fmt.Print("Assignment Name: ")
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        assignName = scanner.Text()
    }

    fmt.Print("Source File Description: ")
    scanner = bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        description = scanner.Text()
    }

    // get class
    fmt.Print("Course: ")
    scanner = bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        class = scanner.Text()
    }

    fmt.Print("Professor: ")
    if scanner.Scan() {
        professor = scanner.Text()
    }

    // assign date / due date
    fmt.Print("Date Assigned: ")
    if scanner.Scan() {
        assignDate = scanner.Text()
    }

    fmt.Print("Date Due: ")
    if scanner.Scan() {
        dueDate = scanner.Text()
    }
    
    // pull student info from conf file

    var student map[string]string
    err = json.Unmarshal(confFile, &student)

    if err != nil {
      fmt.Println(err)
      return
    }

    header := fmt.Sprintf(`#*******************************************************************************
#
#      filename:  %s
#
#   description:  %s
#
#        author:  %s, %s
# AMU e-mail id:  %s
#
#        course:  %s
#    instructor:  %s
#    assignment:  %s
#
#      assigned:  %s
#           due:  %s
#
#*******************************************************************************
`, args[0], description, student["lastName"], student["firstName"], student["email"], class, professor, assignName, assignDate, dueDate)

    newFile, err := os.Create("temp.py")
    if err != nil {
      fmt.Println(err)
      return
    }

    defer newFile.Close()
    
    oldFile, err := os.Open(args[0])
    if err != nil {
      fmt.Println(err)
      return
    }

    defer oldFile.Close()

    _, err = newFile.WriteString(header)
    if err != nil {
      fmt.Println(err)
      return
    }

    scanner = bufio.NewScanner(oldFile)

    // read the file to be appended to and output all of it
    for scanner.Scan() {

        _, err = newFile.WriteString(scanner.Text())
        _, err = newFile.WriteString("\n")
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    newFile.Sync()

    err = os.Rename("temp.py", args[0])
    if err != nil {
      fmt.Println(err)
      return
    }

    fmt.Println("\nCompleted.")
	},
}

func init() {
	rootCmd.AddCommand(addheaderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addheaderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addheaderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
