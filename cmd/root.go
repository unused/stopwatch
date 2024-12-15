package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var srcFile string

// used by some commands to format the output
var format string

// used by some commands for filtering by tags
var filter string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stpw",
	Short: "A cli time tracking tool",
	Long: `Stopwatch is a cli time tracking tool that focus on being transparent
and flexible. Your time entries are stored in a plain text file that you can
edit anytime, adapt and even enrich with comments.`,
	// Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags
// appropriately. This is called by main.main(). It only needs to happen once
// to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Register flags that must be available to all commands. The source file is
// available to all commands via srcFile.
func init() {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	cobra.OnInitialize(createSrcFileIfNotExists)
	defaultFilename := dir + "/.stpw.txt"
	rootCmd.PersistentFlags().StringVar(&srcFile, "file", defaultFilename,
		"source file (default is $HOME/.stpw.txt)")
}

// Create the source file if it does not exist
func createSrcFileIfNotExists() {
	if _, err := os.Stat(srcFile); os.IsNotExist(err) {
		_, err := os.Create(srcFile)
		if err != nil {
			panic(err)
		}
	}
}
