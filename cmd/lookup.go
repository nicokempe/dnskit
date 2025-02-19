package cmd

import (
	"fmt"
	"github.com/nicokempe/dnskit/pkg/dnsutils"
	"github.com/spf13/cobra"
)

var recordType string

var lookupCmd = &cobra.Command{
	Use:   "lookup [hostname]",
	Short: "Perform basic DNS lookups",
	Args:  cobra.ExactArgs(1), // We expect exactly one hostname argument
	RunE: func(cmd *cobra.Command, args []string) error {
		hostname := args[0]
		results, err := dnsutils.Lookup(hostname, recordType)
		if err != nil {
			return err
		}

		for _, r := range results {
			fmt.Println(r)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(lookupCmd)
	lookupCmd.Flags().StringVarP(&recordType, "type", "t", "A", "DNS record type (A, AAAA, CNAME, MX, TXT, etc.)")
}
