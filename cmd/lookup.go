package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/nicokempe/dnskit/pkg/dnsutils"
	"github.com/nicokempe/dnskit/pkg/output"
	"github.com/spf13/cobra"
)

var recordType string

var lookupCmd = &cobra.Command{
	Use:   "lookup [hostname]",
	Short: "Perform basic DNS lookups",
	Args:  cobra.ExactArgs(1), // We expect exactly one hostname argument
	RunE: func(cmd *cobra.Command, args []string) error {
		hostname := args[0]
		results, err := dnsutils.Lookup(hostname, recordType, resolver)
		if err != nil {
			return err
		}

		if outputJSON {
			data := map[string]interface{}{"hostname": hostname, "type": recordType, "records": results}
			b, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(b))
		} else {
			for _, r := range results {
				output.Success(r)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(lookupCmd)
	lookupCmd.Flags().StringVarP(&recordType, "type", "t", "A", "DNS record type (A, AAAA, CNAME, MX, TXT, etc.)")
}
