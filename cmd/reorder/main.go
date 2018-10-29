package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/as27/reorder/pkg/reorder"
	"github.com/as27/reorder/pkg/reorder/fs"
)

const version = "0.9.2"

var (
	flagGap     = flag.Int("gap", 10, "define the gap between the order numbers")
	flagSize    = flag.Int("size", 3, "Number of the digits used 000 <- 3")
	flagMinSize = flag.Int("min", 3, "minimum Number of the digits of the names")
	flagVersion = flag.Bool("version", false, "prints the version")
)

var args []string

func init() {
	flag.Parse()
	args = flag.Args()
}
func main() {
	if *flagVersion {
		fmt.Printf("reorder version: %s\n", version)
		os.Exit(0)
	}
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error getting working directory")
		os.Exit(1)
	}

	// the current wd can be changed with the first arg
	if len(args) > 0 {
		wd, err = filepath.Abs(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "error getting the abs path of ", os.Args[1])
			os.Exit(1)
		}
	}
	_, err = os.Stat(wd)
	if os.IsNotExist(err) {
		fmt.Fprintln(os.Stderr, "the given path does not exist: ", wd)
		os.Exit(1)
	}
	folders := fs.NewFiler(wd, fs.FolderMode)
	log.Println("Reordering folders")
	reorder.Run(folders, *flagGap, *flagSize, *flagMinSize)
	files := fs.NewFiler(wd, fs.FileMode)
	log.Println("Reordering files")
	reorder.Run(files, *flagGap, *flagSize, *flagMinSize)
}
