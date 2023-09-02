/*
Copyright © 2023 TJ Gillis <tj@doublebacklabs.com>
*/
package main

import (
	_ "embed"

	"github.com/Doubleback-Labs/noter/cmd"
)

//go:embed VERSION
var version string

func main() {
	cmd.Execute(version)
}
