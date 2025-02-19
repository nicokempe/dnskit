package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dnskit",
	Short: "DNSKit - A modern DNS analysis CLI",
	Long: `DNSKit is a modern and structured command-line tool 
designed for DNS analysis, penetration testing, and system administration.`,
	// This runs when the user calls `dnskit` without any subcommand
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to DNSKit! Type --help or -h to see available commands.")
		fmt.Println("DNSKit is a modern and structured command-line tool designed for DNS analysis, penetration testing, and system administration.")
	},
}

// Execute is called by main.go
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Here we can define global flags if needed
	// e.g. rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
}
