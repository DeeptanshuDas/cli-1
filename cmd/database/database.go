package database

import (
	"errors"

	"github.com/spf13/cobra"
)

var dbSize, dbName, dbNetworkID, software, softwareVersion, firewallID string
var replicas, numSnapshots int
var publicIP bool

// DBCmd is the root command for the db subcommand
var DBCmd = &cobra.Command{
	Use:     "db",
	Aliases: []string{"database"},
	Short:   "Manage Civo Database ",
	Long:    `Create, update, delete, and list Civo Database.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := cmd.Help()
		if err != nil {
			return err
		}
		return errors.New("a valid subcommand is required")
	},
}

func init() {
	DBCmd.AddCommand(dbCreateCmd)
	DBCmd.AddCommand(dbDeleteCmd)
	// DBCmd.AddCommand(dbUpdateCmd)

	dbCreateCmd.Flags().StringVarP(&softwareVersion, "software-version", "v", "", "the version of the software to use in the creation")
	dbCreateCmd.Flags().IntVarP(&replicas, "replicas", "r", 0, "the number of replicas in addition to the primary node")
	dbCreateCmd.Flags().IntVarP(&numSnapshots, "num-snapshots", "", 1, "the number of snapshots to keep")
	dbCreateCmd.Flags().BoolVarP(&publicIP, "public-ip", "p", true, "whether to assign a public IP to the database")
	dbCreateCmd.Flags().StringVarP(&firewallID, "firewall", "", "", "the firewall to use for the database")
}