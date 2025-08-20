package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/nicokempe/dnskit/pkg/dnsutils"
	"github.com/nicokempe/dnskit/pkg/output"
	"github.com/spf13/cobra"
)

// reverseCmd performs reverse DNS lookups.
var reverseCmd = &cobra.Command{
	Use:   "reverse [ip]",
	Short: "Reverse DNS lookup",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ipAddress := args[0]
		hostnames, err := dnsutils.ReverseLookup(ipAddress, resolver)
		if err != nil {
			return err
		}

		if outputJSON {
			outputData := map[string]interface{}{"ip": ipAddress, "hosts": hostnames}
			jsonOutput, err := json.MarshalIndent(outputData, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(jsonOutput))
		} else {
			for _, hostname := range hostnames {
				output.Success(hostname)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(reverseCmd)
}
