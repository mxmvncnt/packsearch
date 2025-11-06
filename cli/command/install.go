package command

import (
	"flag"
)

var distros = []string{"arch", "ubuntu", "fedora"}

func Install() {
	installFlags := flag.NewFlagSet("install", flag.ExitOnError)
	var distro = ""
	installFlags.StringVar(&distro, "distro", "", "Parameter value")
	installFlags.StringVar(&distro, "d", "", "Parameter value")

	if distro == "" {

	}

	// TODO: if distro param is not set, assume the package name is the one of the current distro
	// For that, use the command "lsb_release -i -s"

}
