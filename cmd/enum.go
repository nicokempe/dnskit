package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/nicokempe/dnskit/pkg/dnsutils"
	"github.com/nicokempe/dnskit/pkg/output"
	"github.com/spf13/cobra"
)

var (
	// enumWordlist is the path to the subdomain wordlist.
	enumWordlist string
	// enumConcurrency controls the number of concurrent lookups.
	enumConcurrency int
)

// enumCmd enumerates subdomains for a domain.
var enumCmd = &cobra.Command{
	Use:   "enum [domain]",
	Short: "Enumerate subdomains",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		domain := args[0]
		subdomains, err := dnsutils.Enumerate(domain, enumWordlist, enumConcurrency, resolver)
		if err != nil {
			return err
		}

		if outputJSON {
			outputData := map[string]interface{}{"domain": domain, "subdomains": subdomains}
			jsonOutput, err := json.MarshalIndent(outputData, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(jsonOutput))
		} else {
			for _, subdomain := range subdomains {
				output.Success(subdomain)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(enumCmd)
	enumCmd.Flags().StringVarP(&enumWordlist, "wordlist", "w", "", "Path to subdomain wordlist")
	enumCmd.Flags().IntVarP(&enumConcurrency, "concurrency", "c", 10, "Number of concurrent lookups")
}
