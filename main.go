package main

import (
	"flag"
	"fmt"
	"io"
	"os"
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
	flag.StringVar(&configPath, "config", usr.HomeDir+"/.gorespect.json", "Path to config file")
	flag.Parse()

	cs := NewConfigStorage(configPath)
	config := cs.Load()
	defer cs.Save(config)

	dir, err = filepath.Abs(dir)
	if err != nil {
		fmt.Printf("Can not get absolute path for dir: %s. Error: %e\n", dir, err)
		os.Exit(1)
	}

	err = setUpConfig(config, os.Stdout, os.Stdin)
	if err != nil {
		fmt.Printf("Could not set up config: %s\n", err)
		os.Exit(1)
	}

	github := &GithubRespecter{
		Username: config.Github.Username,
		Token:    config.Github.Token,
		Out:      os.Stdout,
		In:       os.Stdin,
	}

	packages, err := getImports(dir)
	if err != nil {
		fmt.Printf("Can not get dependencies: %s", err)
		os.Exit(1)
	}

	packages = github.FilterRespectable(packages)
	if len(packages) == 0 {
		fmt.Println("We didn't find any dependency")
		os.Exit(0)
	}

	maxPackageLength := maxStringLength(packages)

	for _, p := range packages {
		fmt.Print(padRight(p, " ", maxPackageLength) + " — ")
		err := github.SayRespect(p)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Println("Respected")
		}
	}
}

func setUpConfig(config *Config, out io.Writer, in io.Reader) error {
	var err error

	if config.Github.Username == "" {
		config.Github.Username, err = promptGithubUsername(out, in)
		if err != nil {
			return err
		}
	}

	if config.Github.Token == "" {
		config.Github.Token, err = promptGithubToken(out, in)
		if err != nil {
			return err
		}
	}

	return nil
}
