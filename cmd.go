package main

import (
	"flag"
	"fmt"
	"os"
)

var BinVersion string

func cmdExec() {
	flag.Usage = func() {
		fmt.Printf("Name:\n")
		fmt.Printf("    Template for a web app cmd.\n")
		fmt.Printf("Usage of sled:\n")
		fmt.Printf("    sled <opts> <flags>\n")
		fmt.Printf("Commands:\n")
		fmt.Printf("    version   Shows the version information.\n")
		fmt.Printf("\n")
		fmt.Printf("Options:\n")
		fmt.Printf("    -h --help     Show this screen.\n")
	}

	flag.Parse()

	if len(os.Args) < 2 {
		return
	}

	switch os.Args[1] {
	case "version":
		fmt.Printf("Made with love.\n")
		fmt.Printf("Version: v%s\n", BinVersion)
		fmt.Printf("License: MIT\n")
		os.Exit(0)
	}

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

}
