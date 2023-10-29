package cmd

import (
	"fmt"

	"github.com/RATIU5/imgsrt/internal"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "output the version of imgsrt",
	Long:  "all software has versions. see the current version of imgsrt",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("imgsrt v%s\n", internal.Version)
	},
}
