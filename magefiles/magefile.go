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

func Mac() error {
	os.Setenv("GOOS", "darwin")
	os.Setenv("GOARCH", "amd64")
	return build("mac")
}

func Linux() error {
	os.Setenv("GOOS", "linux")
	os.Setenv("GOARCH", "amd64")
	return build("linux")
}

func Windows() error {
	os.Setenv("GOOS", "windows")
	os.Setenv("GOARCH", "amd64")
	return build("windows")
}

func build(target string) error {
	mg.Deps(InstallDeps)
	fmt.Println("Building...")
	name := fmt.Sprintf("./bin/%s", appName)
	switch target {
	case "mac":
		name = fmt.Sprintf("%s-darwin_amd64", name)
	case "linux":
		name = fmt.Sprintf("%s-linux_amd64", name)
	case "windows":
		name = fmt.Sprintf("%s-windows_amd64.exe", name)
	}

	cmd := exec.Command("go", "build", "-o", name, ".")
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
	home, _ := os.UserHomeDir()
	os.RemoveAll(fmt.Sprintf("%s/.noter", home))
	os.RemoveAll("bin/")
	os.RemoveAll(appName)
}
