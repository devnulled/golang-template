// Build is the build pipeline for this repository.
package main

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/boot"
)

// Directories used in repository, relative to this program's working
// directory. goyek.sh and goyek.ps1 both cd into build/ before `go run .`,
// so the build module is "." and the repository root is one level up.
const (
	dirRoot  = ".."
	dirBuild = "."
)

func main() {
	goyek.SetDefault(all)
	boot.Main()
}
