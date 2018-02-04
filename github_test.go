package main

import "testing"

func TestGithubRespecter_CanProcess(t *testing.T) {
	testCases := []map[string]interface{} {
		{
			"str": "",
			"res": false,
		},
		{
			"str": "os",
			"res": false,
		},
		{
			"str": "package/github.com/com",
			"res": false,
		},
		{
			"str": "github.com/iamsalnikov/my-respect",
			"res": true,
		},
	}

	g := GithubRespecter{}
	for _, testCase := range testCases {
		res := g.CanProcess(testCase["str"].(string))
		expected := testCase["res"].(bool)

		if res != expected {
			t.Errorf("I expected to get %v but got %v", expected, res)
		}
	}
}

func TestGithubRespecter_normalizePackageName(t *testing.T) {
	testCases := []map[string]interface{} {
		{
			"package": "",
			"useHost": false,
			"out": "",
		},
		{
			"package": "github.com",
			"useHost": false,
			"out": "github.com",
		},
		{
			"package": "github.com/iamsalnikov",
			"useHost": false,
			"out": "github.com/iamsalnikov",
		},
		{
			"package": "github.com/iamsalnikov/my-respect",
			"useHost": false,
			"out": "iamsalnikov/my-respect",
		},
		{
			"package": "github.com/iamsalnikov/my-respect/subpackage",
			"useHost": false,
			"out": "iamsalnikov/my-respect",
		},
		{
			"package": "github.com/iamsalnikov/my-respect/subpackage",
			"useHost": true,
			"out": "github.com/iamsalnikov/my-respect",
		},
	}

	g := GithubRespecter{}
	for _, testCase := range testCases {
		res := g.normalizePackageName(testCase["package"].(string), testCase["useHost"].(bool))
		expected := testCase["out"].(string)

		if res != expected {
			t.Errorf("I expected to get %s but got %s", expected, res)
		}
	}
}

func TestGithubRespecter_FilterRespectable(t *testing.T) {
	testCases := []map[string]interface{}{
		{
			"pkgs": []string{},
			"ctrl": map[string]bool{},
		},
		{
			"pkgs": []string{"os", "bitbucket.org/some/package"},
			"ctrl": map[string]bool{},
		},
		{
			"pkgs": []string{"os", "github.com/iamsalnikov/my-respect", "github.com/some/package/ping"},
			"ctrl": map[string]bool{
				"github.com/iamsalnikov/my-respect": true,
				"github.com/some/package": true,
			},
		},
	}

	g := GithubRespecter{}
	for _, testCase := range testCases {
		res := g.FilterRespectable(testCase["pkgs"].([]string))
		ctrl := testCase["ctrl"].(map[string]bool)

		if len(res) != len(ctrl) {
			t.Errorf("I got different lengths (%d and %d)", len(res), len(ctrl))
			continue
		}

		if len(ctrl) == 0 {
			continue
		}

		for _, r := range res {
			_, ok := ctrl[r]
			if !ok {
				t.Errorf("I expected to see package %s but do not see it", r)
			}
		}
	}
}