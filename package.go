package main

import (
	"os"
	"strings"

	"github.com/golang/dep/gps/pkgtree"
)

func getImports(dir string) ([]string, error) {
	currentPackage := strings.Replace(dir, os.Getenv("GOPATH")+"/src/", "", 1)

	pkgs, err := pkgtree.ListPackages(
		dir,
		"",
	)

	if err != nil {
		return []string{}, err
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

	return imports, nil
}
