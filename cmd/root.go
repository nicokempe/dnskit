package cmd

import (
	"fmt"
	"github.com/nicokempe/dnskit/pkg/version"
	"github.com/spf13/cobra"
)

var (
	// outputJSON controls whether commands should emit JSON instead of
	// colorized text. It is exposed as a persistent flag on the root
	// command and therefore available to all subcommands.
	outputJSON bool

	// resolver allows the user to override the system DNS resolver. It is
	// a hostname or IP (optionally including a port) that will be used for
	// all DNS queries when provided.
	resolver string
)

// rootCmd is the base command for DNSKit.
var rootCmd = &cobra.Command{
	Use:   "dnskit",
	Short: "DNSKit - A modern DNS analysis CLI",
	Long: `DNSKit is a modern and structured command-line tool
designed for DNS analysis, penetration testing, and system administration.`,
	Version: version.Version,
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
	// Global persistent flags shared by all subcommands
	rootCmd.PersistentFlags().BoolVar(&outputJSON, "json", false, "Output results in JSON format")
	rootCmd.PersistentFlags().StringVar(&resolver, "resolver", "", "Use a custom DNS resolver (ip[:port])")
}
