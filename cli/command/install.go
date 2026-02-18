package command

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"strings"
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
		distro = detectCurrentDistro()
		packagesSliceIndex = 2
	}

	if !slices.Contains(distros, distro) {
		fmt.Printf("Invalid distro: %s\n", distro)
		return
	}

	programNames := os.Args[packagesSliceIndex:]

	fmt.Printf("install cmd: %s\n", Install(distro, programNames))
}

func detectCurrentDistro() string {
	// alternatively could look into using lsb_release command
	content, _ := os.ReadFile("/etc/os-release")
	for _, line := range strings.Split(string(content), "\n") {
		if strings.HasPrefix(line, "ID=") {
			return strings.TrimPrefix(line, "ID=")
		}
	}

	return "unknown"
}
