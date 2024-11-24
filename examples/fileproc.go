package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/vh8t/argparser"
)

func main() {
	parser := argparser.NewRule("fileproc", "File Processor", "v1.0", true).
		AddStringFlag("file", "f", "Path to the input file", true).
		AddStringFlag("mode", "m", "Processing mode: 'content' or 'summary'", false, "content").
		AddBoolFlag("verbose", "v", "Enable verbose output")

	err := parser.Parse(os.Args[1:])
	help := parser.GetBoolFlag("help")

	if err != nil && !help {
		fmt.Println(err)
		fmt.Println(parser.Help())
		return
	} else if help {
		fmt.Println(parser.Help())
		return
	}

	file, _ := parser.GetStringFlag("file")
	mode, _ := parser.GetStringFlag("mode")
	verbose := parser.GetBoolFlag("verbose")

	if verbose {
		fmt.Printf("[INFO] Using file: %s\n", file)
		fmt.Printf("[INFO] Mode: %s\n", mode)
	}

	switch strings.ToLower(mode) {
	case "content":
		fmt.Printf("Processing file '%s'...\n", file)
		fmt.Println("File content: (simulated data)")
		fmt.Println("Lorem ipsum dolor sit amet...")
	case "summary":
		fmt.Printf("Processing file '%s' in summary mode...\n", file)
		fmt.Println("Summary: (simulated data)")
		fmt.Println("Lines: 10, Words: 100, Characters: 500")
	default:
		fmt.Printf("Error: Unknown mode '%s'. Use --help for available options.\n", mode)
	}
}
