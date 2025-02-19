package cmd

import (
	"fmt"
	"github.com/nicokempe/dnskit/pkg/dnsutils"
	"github.com/spf13/cobra"
)

var enumWordlist string

var enumCmd = &cobra.Command{
	Use:   "enum [domain]",
	Short: "Enumerate subdomains",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		domain := args[0]
		subdomains, err := dnsutils.Enumerate(domain, enumWordlist)
		if err != nil {
			return err
		}

		for _, sub := range subdomains {
			fmt.Println(sub)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(enumCmd)
	enumCmd.Flags().StringVarP(&enumWordlist, "wordlist", "w", "", "Path to subdomain wordlist")
}
