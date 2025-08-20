package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/nicokempe/dnskit/pkg/dnsutils"
	"github.com/nicokempe/dnskit/pkg/output"
	"github.com/spf13/cobra"
)

var reverseCmd = &cobra.Command{
	Use:   "reverse [ip]",
	Short: "Reverse DNS lookup",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ip := args[0]
		hosts, err := dnsutils.ReverseLookup(ip, resolver)
		if err != nil {
			return err
		}

		if outputJSON {
			data := map[string]interface{}{"ip": ip, "hosts": hosts}
			b, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(b))
		} else {
			for _, h := range hosts {
				output.Success(h)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(reverseCmd)
}
