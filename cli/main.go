package main

import (
	"fmt"
	"os"

	"github.com/mxmvncnt/packsearch/cli/command"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	subcommand := os.Args[1]
	switch subcommand {
	case "install":
		command.InstallCommand()
	case "help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", subcommand)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: pck <command> [options]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  install    InstallCommand a package")
	fmt.Println("  help       Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  pck install --param value")
	fmt.Println("  pck install --param value --verbose package1 package2")
}
