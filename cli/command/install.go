package command

import (
	"flag"
	"fmt"
	"os"
)

var distros = []string{"arch", "ubuntu", "fedora"}

func InstallCommand() {
	installFlags := flag.NewFlagSet("install", flag.ExitOnError)
	var distro = ""
	installFlags.StringVar(&distro, "distro", "", "Parameter value")
	installFlags.StringVar(&distro, "d", "", "Parameter value")

	err := installFlags.Parse(os.Args[2:])
	if err != nil {
		return
	}

	if distro == "" {
		// TODO: if distro param is not set, assume the package name is the one of the current distro
		// For that, use the command "lsb_release -i -s" or find a way to get the info directly
	}

	fmt.Printf("install cmd: %s\n", Install(distro, []string{"test", "test2"}))
}
