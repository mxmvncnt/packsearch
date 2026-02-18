package command

import (
	"fmt"
	"strings"

	"github.com/mxmvncnt/packsearch/cli/util"
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
	if !isCommandRunAsRoot() {
		util.ExitWithError("permission denied: you must run this command as root")
	}
	packagesStr := strings.Join(packages, " ")

	fromPackageManager := packageManager[distro]
	toInstallCommand := installCommand[fromPackageManager]

	return fmt.Sprintf("%s %s %s", fromPackageManager, toInstallCommand, packagesStr)
}

// function disabled temporarly
// will need to detect wether the package manager requires root (some install in user profile, like Flathub or Homebrew)
func isCommandRunAsRoot() bool {
	//currentUser, err := user.Current()
	//if err != nil {
	//	log.Fatalf("could not get current user: %s", err)
	//}
	//return currentUser.Username == "root"
	return true
}
