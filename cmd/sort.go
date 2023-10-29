package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/RATIU5/imgsrt/internal/utils"
	"github.com/spf13/cobra"
)

var (
	enteredOutdir string
	enteredIndir  string
	pattern       string
	isVerbose     bool

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
	cmdSort.Flags().StringVarP(&enteredIndir, "input", "i", "", "assign the input path of media")
	cmdSort.Flags().StringVarP(&enteredOutdir, "output", "o", "", "assign the output path of media")
	cmdSort.Flags().StringVarP(&pattern, "pattern", "p", "", "set the sorting pattern;\ny - year\nm - month")
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
		return
	}

	err = handleAssignedPath(&outdir, enteredOutdir, filepath.Join(cwd, "output"))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = handleAssignedPath(&indir, enteredIndir, cwd)
	if err != nil {
		fmt.Println(err)
		return
	}

	if pattern == "" {
		pat = "y/m"
	} else {
		pat = pattern
	}

	dirGen := parsePattern(pat)
}

func handleAssignedPath(target *string, assigned string, def string) error {
	if assigned == "" {
		res, err := utils.NormalizePath(def)
		if err != nil {
			return err
		}
		*target = res
	} else {
		res, err := utils.NormalizePath(*target)
		if err != nil {
			return err
		}
		*target = res
	}
	return nil
}

func parsePattern(pattern string) func(time.Time) string {
	parts := strings.Split(pattern, "/")
	return func(t time.Time) string {
		var dirs []string
		for _, part := range parts {
			switch part {
			case "y":
				dirs = append(dirs, fmt.Sprintf("%d", t.Year()))
			case "m":
				dirs = append(dirs, fmt.Sprintf("%02d", t.Month()))
			// ... handle other pattern parts as needed
			default:
				dirs = append(dirs, part) // treat unrecognized parts as literal directory names
			}
		}
		return strings.Join(dirs, string(os.PathSeparator))
	}
}
