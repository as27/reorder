package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "0.1.0"

var (
	flagGap     = flag.Int("gap", 10, "define the gap between the order numbers")
	flagSize    = flag.Int("size", 3, "Number of the digits used 000 <- 3")
	flagVersion = flag.Bool("version", false, "prints the version")
)

func main() {
	flag.Parse()
	if *flagVersion {
		fmt.Printf("reorder version: %s\n", version)
	}
	fmt.Printf("%#v\n", os.Args)
}
