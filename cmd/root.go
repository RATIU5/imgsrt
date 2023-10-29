package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "imgsrt",
	Short: "imgsrt is a very fast image sorter",
	Long:  "imgsrt is a very fast image sorter with flexibility",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
