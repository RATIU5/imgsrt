package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/RATIU5/imgsrt/internal/utils"
	"github.com/rwcarlsen/goexif/exif"
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
		pat = "y/M"
	} else {
		pat = pattern
	}

	dirGen := parsePattern(pat)
	fmt.Printf("parsing with pattern: '%s', in directory '%s', outputting in '%s'", pat, indir, outdir)
	err = processFiles(indir, outdir, dirGen)
	if err != nil {
		fmt.Println(err)
	}
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
			case "M":
				dirs = append(dirs, fmt.Sprintf("%s", t.Month().String()))
			case "m":
				dirs = append(dirs, fmt.Sprintf("%02d", t.Month()))
			case "d":
				dirs = append(dirs, fmt.Sprintf("%02d", t.Day()))
			default:
				dirs = append(dirs, part) // treat unrecognized parts as literal directory names
			}
		}
		return strings.Join(dirs, string(os.PathSeparator))
	}
}

// DetermineDestinationDir determines the destination directory based on file properties.
func DetermineDestinationDir(info os.FileInfo, path string, outdir string, dirGen func(time.Time) string) (string, error) {
	destDir := filepath.Join(outdir, "others")
	if isImageOrVideo(info.Name()) {
		date, err := getFileCreationDate(path)
		if err == nil {
			dirString := dirGen(date)
			destDir = filepath.Join(outdir, dirString)
		}
	}
	if info.ModTime().IsZero() {
		destDir = filepath.Join(outdir, "unknown")
	}
	return destDir, nil
}

// getFileCreationDate extracts the creation date from file metadata.
func getFileCreationDate(path string) (time.Time, error) {
	f, err := os.Open(path)
	if err != nil {
		return time.Time{}, err
	}
	defer f.Close()

	x, err := exif.Decode(f)
	if err != nil {
		return time.Time{}, err
	}

	date, err := x.DateTime()
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}

// processFile processes a single file.
func processFile(path string, destDir string) error {
	err := utils.EnsureDir(destDir)
	if err != nil {
		return err
	}
	if filepath.Base(path) != "imgsrt.exe" {
		destPath := filepath.Join(destDir, filepath.Base(path))
		err = os.Rename(path, destPath) // use io.Copy if you want to copy instead of move
	}
	return err
}

// processFiles processes all files under indir.
func processFiles(indir string, outdir string, dirGen func(time.Time) string) error {
	return filepath.Walk(indir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}

		if !info.IsDir() {
			destDir, err := DetermineDestinationDir(info, path, outdir, dirGen)
			if err != nil {
				fmt.Printf("error: %s\n", err)
			}
			err = processFile(path, destDir)
			if err != nil {
				fmt.Printf("error: %v\n", err)
			}
		}

		return nil
	})
}

var imageExtensions = []string{
	".jpg", ".jpeg", ".png", ".gif", ".bmp",
	".tiff", ".webp", ".raw", ".cr2", ".nef",
}

var videoExtensions = []string{
	".mp4", ".mkv", ".flv", ".avi", ".mov",
	".wmv",
}

func isImageOrVideo(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	for _, imageExt := range imageExtensions {
		if strings.EqualFold(ext, imageExt) {
			return true
		}
	}
	for _, videoExt := range videoExtensions {
		if strings.EqualFold(ext, videoExt) {
			return true
		}
	}
	return false
}
