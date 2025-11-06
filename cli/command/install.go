package command

import (
	"flag"
	"fmt"
	"os"
	"slices"
)

var distros = []string{"arch", "ubuntu", "fedora"}

func InstallCommand() {
	installFlags := flag.NewFlagSet("install", flag.ExitOnError)
	var distro = ""
	installFlags.StringVar(&distro, "distro", "", "Parameter value")
	installFlags.StringVar(&distro, "d", "", "Parameter value")

	err := installFlags.Parse(os.Args[2:4])
	if err != nil {
		return
	}

	// packagesSliceIndex is for when the "Distro" parameter is not provided, we trim the Args at the right place
	// since it will be 2 args shorter than expected:
	//   <pck> install --distro fedora pkgName1 pgkName2 <-- 2:4 are --distro and fedora, so the packages start at 4
	//   <pck> install pkgName1 pkgName2 <-- 2:4 are pgkName1 and pkgName2, so the packages start at 2
	var packagesSliceIndex = 4
	if distro == "" {
		distro = "ubuntu"
		packagesSliceIndex = 2
		// TODO: if distro param is not set, assume the package name is the one of the current distro
		// For that, use the command "lsb_release -i -s" or find a way to get the info directly
	}

	if !slices.Contains(distros, distro) {
		fmt.Printf("Invalid distro: %s\n", distro)
		return
	}

	programNames := os.Args[packagesSliceIndex:]

	fmt.Printf("install cmd: %s\n", Install(distro, programNames))
}
