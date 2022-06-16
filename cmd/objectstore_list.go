package cmd

import (
	"os"

	"github.com/civo/cli/config"
	"github.com/civo/cli/utility"
	"github.com/spf13/cobra"
)

var objectStoreListCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list", "all"},
	Example: `civo objectstore ls`,
	Short:   "List all objectstores",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := config.CivoAPIClient()
		if err != nil {
			utility.Error("Creating the connection to Civo's API failed with %s", err)
			os.Exit(1)
		}

		objectStores, err := client.ListObjectStores()
		if err != nil {
			utility.Error("%s", err)
			os.Exit(1)
		}

		ow := utility.NewOutputWriter()
		for _, objectStore := range objectStores {
			ow.StartLine()
			ow.AppendDataWithLabel("id", objectStore.ID, "ID")
			ow.AppendDataWithLabel("generated_name", objectStore.GeneratedName, "Generated Name")
			ow.AppendDataWithLabel("size", objectStore.MaxSize, "Size")
			ow.AppendDataWithLabel("bucket_url", objectStore.BucketURL, "Object Store Endpoint")
			ow.AppendDataWithLabel("s3_region", "default", "S3 Region")
			ow.AppendDataWithLabel("status", objectStore.Status, "Status")
		}

		switch outputFormat {
		case "json":
			ow.WriteMultipleObjectsJSON(prettySet)
		case "custom":
			ow.WriteCustomOutput(outputFields)
		default:
			ow.WriteTable()
		}
	},
}
