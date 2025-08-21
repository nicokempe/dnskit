package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/nicokempe/dnskit/pkg/dnsutils"
	"github.com/nicokempe/dnskit/pkg/output"
	"github.com/spf13/cobra"
)

// nameserverAddr is the target nameserver for zone transfers.
var nameserverAddr string

// transferCmd attempts DNS zone transfers (AXFR).
var transferCmd = &cobra.Command{
	Use:   "transfer [domain]",
	Short: "Attempt a DNS zone transfer (AXFR)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		domain := args[0]
		zoneRecords, err := dnsutils.ZoneTransfer(domain, nameserverAddr)
		if err != nil {
			fmt.Printf("Zone transfer failed: %v\n", err)
			return nil
		}

		if outputJSON {
			outputData := map[string]interface{}{"domain": domain, "records": zoneRecords}
			jsonOutput, err := json.MarshalIndent(outputData, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(jsonOutput))
		} else {
			for _, zoneRecord := range zoneRecords {
				output.Success(zoneRecord)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(transferCmd)
	transferCmd.Flags().StringVarP(&nameserverAddr, "nameserver", "n", "", "Nameserver to use for zone transfer")
}
