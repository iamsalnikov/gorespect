package main

import (
	"os"
	"strings"

	"github.com/golang/dep/gps/pkgtree"
)

func getImports(dir string) ([]string, error) {
	currentPackage := strings.Replace(dir, os.Getenv("GOPATH")+"/src/", "", 1)

	pkgs, err := pkgtree.ListPackages(dir, "")

	if err != nil {
		return []string{}, err
	}

	importsMap := make(map[string]bool)
	imports := make([]string, 0)
	for _, p := range pkgs.Packages {
		for _, pn := range p.P.Imports {
			if strings.Index(pn, currentPackage) == 0 {
				continue
			}

			if !importsMap[pn] {
				importsMap[pn] = true
				imports = append(imports, pn)
			}
		}
	}

	return imports, nil
}
