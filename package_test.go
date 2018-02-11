package main

import "testing"

func TestGetImports_NoDir(t *testing.T) {
	dir := "testdata/unexisting_dir"
	imports, err := getImports(dir)

	if err == nil {
		t.Errorf("I expected to get error, but got nil")
	}

	if len(imports) > 0 {
		t.Errorf("I expected to get empty imports, but current length is %d", len(imports))
	}
}

func TestGetImports_EmptyDir(t *testing.T) {
	dir := "testdata/src/empty"
	imports, err := getImports(dir)

	if err != nil {
		t.Errorf("I got unexpected error: %s", err)
	}

	if len(imports) > 0 {
		t.Errorf("I expected to get empty imports, but current length is %d", len(imports))
	}
}

func TestGetImports_GoodDir(t *testing.T) {
	dir := "testdata/src/good"
	imports, err := getImports(dir)

	if err != nil {
		t.Errorf("I got unexpected error: %s", err)
	}

	expected := map[string]bool{
		"github.com/golang/dep":     true,
		"github.com/golang/dep/gps": true,
		"os":   true,
		"fmt":  true,
		"sync": true,
	}

	if len(expected) != len(imports) {
		t.Errorf("I expected to get %d packages but got %d", len(expected), len(imports))
	}

	for _, pn := range imports {
		if !expected[pn] {
			t.Errorf("I expected to get package %s, but do not see it", pn)
		}
	}
}
