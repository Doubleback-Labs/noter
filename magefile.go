//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)

var appName = "noter"

func Build() error {
	mg.Deps(InstallDeps)
	fmt.Println("Building...")
	cmd := exec.Command("go", "build", "-o", appName, ".")
	return cmd.Run()
}

func InstallDeps() error {
	fmt.Println("Installing Deps...")
	cmd := exec.Command("/bin/sh", "-c", "go get github.com/spf13/viper; go get github.com/spf13/cobra; go get github.com/rs/zerolog/log;")
	return cmd.Run()
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll(appName)
}
