package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/nicokempe/dnskit/pkg/dnsutils"
	"github.com/nicokempe/dnskit/pkg/output"
	"github.com/spf13/cobra"
)

var nameserver string

var transferCmd = &cobra.Command{
	Use:   "transfer [domain]",
	Short: "Attempt a DNS zone transfer (AXFR)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		domain := args[0]
		records, err := dnsutils.ZoneTransfer(domain, nameserver)
		if err != nil {
			fmt.Printf("Zone transfer failed: %v\n", err)
			return nil
		}

		if outputJSON {
			data := map[string]interface{}{"domain": domain, "records": records}
			b, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(b))
		} else {
			for _, record := range records {
				output.Success(record)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(transferCmd)
	transferCmd.Flags().StringVarP(&nameserver, "nameserver", "n", "", "Nameserver to use for zone transfer")
}
