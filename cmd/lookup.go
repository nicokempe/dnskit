package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/nicokempe/dnskit/pkg/dnsutils"
	"github.com/nicokempe/dnskit/pkg/output"
	"github.com/spf13/cobra"
)

// recordType specifies the DNS record type for lookup.
var recordType string

// lookupCmd performs DNS record lookups for a given hostname.
var lookupCmd = &cobra.Command{
	Use:   "lookup [hostname]",
	Short: "Perform basic DNS lookups",
	Args:  cobra.ExactArgs(1), // We expect exactly one hostname argument
	RunE: func(cmd *cobra.Command, args []string) error {
		hostname := args[0]
		recordResults, err := dnsutils.Lookup(hostname, recordType, resolver)
		if err != nil {
			return err
		}

		if outputJSON {
			outputData := map[string]interface{}{"hostname": hostname, "type": recordType, "records": recordResults}
			jsonOutput, err := json.MarshalIndent(outputData, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(jsonOutput))
		} else {
			for _, record := range recordResults {
				output.Success(record)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(lookupCmd)
	lookupCmd.Flags().StringVarP(&recordType, "type", "t", "A", "DNS record type (A, AAAA, CNAME, MX, TXT, etc.)")
}
