package command

import (
	"fmt"
	"strings"
)

var packageManager = map[string]string{
	"fedora":    "dnf",
	"redhat":    "dnf",
	"centos":    "dnf",
	"almalinux": "dnf",
	"ubuntu":    "apt",
	"debian":    "apt",
}

var installCommand = map[string]string{
	"dnf": "install -y",
	"apt": "install -y",
}

// Install will run the native package manager to install the package(s) for the distro
func Install(distro string, packages []string) string {
	packagesStr := strings.Join(packages, " ")

	fromPackageManager := packageManager[distro]
	toInstallCommand := installCommand[fromPackageManager]

	return fmt.Sprintf("%s %s %s", fromPackageManager, toInstallCommand, packagesStr)
}
