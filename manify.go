package main

import (
	"flag"
	"fmt"
)

var args struct {
	source string
	target string
}

/**
 * @author Christin Gruber
 */
func main() {
	fmt.Println("Let's run manify...")

	flag.StringVar(&args.source, "source", "source.md", "Set the markdown source file to convert")
	flag.StringVar(&args.target, "target", "target.html", "Set the html target file with converted contents")
	flag.Parse()

	fmt.Printf("with args --source %s and --target %s...", args.source, args.target)
}

func compile() {
	// Compile markdown
}