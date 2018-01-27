package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("can not get dir", err)
	}

	packages := getImports(dir)

	github := &GithubThanker{}

	for _, p := range packages {
		if github.CanProcess(p) {
			github.SayThankYou(p)
		}
	}
}




