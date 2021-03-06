package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
const Version = "1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of nexusSync",
	Long:  `All software has versions. This is nexusSync's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("nexusSync version is v%s\n", Version)
	},
}
