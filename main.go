package main

import (
	"fmt"
	"os"
	"flag"
	"os/user"
	"path/filepath"
)

func main() {
	defaultDir, err := os.Getwd()
	if err != nil {
		fmt.Println("can not get current dir", err)
		os.Exit(1)
	}

	usr, err := user.Current()
	if err != nil {
		fmt.Println("can not get current user", err)
		os.Exit(1)
	}

	var dir string
	var configPath string

	flag.StringVar(&dir, "dir", defaultDir, "Directory with package")
	flag.StringVar(&configPath, "config", usr.HomeDir + "/.thanker.json", "Path to config file")
	flag.Parse()

	config := NewConfig(configPath)
	defer config.Save()

	dir, err = filepath.Abs(dir)
	if err != nil {
		fmt.Printf("Can not get absolute path for dir: %s. Error: %e\n", dir, err)
		os.Exit(1)
	}

	packages := getImports(dir)
	if len(packages) == 0 {
		fmt.Println("We didn't find any dependency")
		os.Exit(0)
	}

	github := &GithubThanker{
		Config: config,
	}

	maxPackageLength := maxStringLength(packages)

	for _, p := range packages {
		fmt.Print(padRight(p, " ", maxPackageLength) + " — ")

		if github.CanProcess(p) {
			err := github.SayThankYou(p)

			if err != nil {
				fmt.Printf("Facepalm (%s)\n", err)
			} else {
				fmt.Println("Starred")
			}
		} else {
			fmt.Println("Can not say thanks")
		}
	}
}




