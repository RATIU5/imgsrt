package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	outputPath string
	inputPath  string
	pattern    string
	isVerbose  bool

	cmdSort = &cobra.Command{
		Use:   "sort",
		Short: "sort images",
		Long:  "sort images based on a configuration",
		Run: func(cmd *cobra.Command, args []string) {
			out := handleSortCommand(cmd, args)
			fmt.Println(out)
		},
	}
)

func init() {
	cmdSort.Flags().StringVarP(&inputPath, "input", "i", "", "assign the input path of media")
	cmdSort.Flags().StringVarP(&outputPath, "output", "o", "", "assign the output path of media")
	cmdSort.Flags().StringVarP(&pattern, "pattern", "p", "", "set the sorting pattern to use")
	cmdSort.Flags().BoolVarP(&isVerbose, "verbose", "v", false, "set verbose output")
	rootCmd.AddCommand(cmdSort)
}

func handleSortCommand(cmd *cobra.Command, args []string) string {
	var output string

	return output
}
