package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/nicokempe/dnskit/pkg/dnsutils"
	"github.com/nicokempe/dnskit/pkg/output"
	"github.com/spf13/cobra"
)

var (
	enumWordlist    string
	enumConcurrency int
)

var enumCmd = &cobra.Command{
	Use:   "enum [domain]",
	Short: "Enumerate subdomains",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		domain := args[0]
		subs, err := dnsutils.Enumerate(domain, enumWordlist, enumConcurrency, resolver)
		if err != nil {
			return err
		}

		if outputJSON {
			data := map[string]interface{}{"domain": domain, "subdomains": subs}
			b, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(b))
		} else {
			for _, sub := range subs {
				output.Success(sub)
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
