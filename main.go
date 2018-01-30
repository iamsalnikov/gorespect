package main

import (
	"fmt"
	"os"
	"flag"
	"os/user"
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

	packages := getImports(dir)

	github := &GithubThanker{
		Config: config,
	}

	for _, p := range packages {
		if github.CanProcess(p) {
			fmt.Print(p + " : ")
			err := github.SayThankYou(p)

			if err != nil {
				fmt.Printf("Facepalm (%s)\n", err)
			} else {
				fmt.Println("Starred")
			}
		}
	}
}




