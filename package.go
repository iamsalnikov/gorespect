package main

import (
	"fmt"
	"strings"
	"os"
	"github.com/golang/dep/gps/pkgtree"
)

func getImports(dir string) []string {
	currentPackage := strings.Replace(dir, os.Getenv("GOPATH") + "/src/", "", 1)

	pkgs, err := pkgtree.ListPackages(
		dir,
		"",
	)

	if err != nil {
		fmt.Println("error: ", err)
		return []string{}
	}

	importsMap := make(map[string]bool)
	for _, p := range pkgs.Packages {
		for _, pn := range p.P.Imports {
			if strings.Index(pn, currentPackage) == 0 {
				continue
			}

			importsMap[pn] = true
		}
	}

	imports := make([]string, len(importsMap))
	i := 0
	for imp := range importsMap {
		imports[i] = imp
		i++
	}

	return imports
}
