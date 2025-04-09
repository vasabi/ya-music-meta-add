package metadata

import "github.com/spf13/cobra"

func init() {
	Metadata.AddCommand(start)
}

var Metadata = &cobra.Command{
	Use:   "metadata",
	Short: "Print tickets methods",
}
