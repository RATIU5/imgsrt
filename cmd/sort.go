package cmd

import (
	"fmt"
	"os"
	"path/filepath"

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
			handleSortCommand(cmd, args)
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

func handleSortCommand(cmd *cobra.Command, args []string) {
	var outdir string
	var indir string
	var pat string
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("error: failed to read cwd")
	}

	if outputPath == "" {
		outdir = filepath.Join(cwd, "output")
	} else {
		outdir = outputPath
	}

	if inputPath == "" {
		indir = cwd
	} else {
		indir = inputPath
	}

	if pattern == "" {
		pat = "y/m"
	} else {
		pat = pattern
	}
}
