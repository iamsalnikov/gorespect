package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGithubRespecter_CanProcess(t *testing.T) {
	testCases := []map[string]interface{}{
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
	testCases := []map[string]interface{}{
		{
			"package": "",
			"useHost": false,
			"out":     "",
		},
		{
			"package": "github.com",
			"useHost": false,
			"out":     "github.com",
		},
		{
			"package": "github.com/iamsalnikov",
			"useHost": false,
			"out":     "github.com/iamsalnikov",
		},
		{
			"package": "github.com/iamsalnikov/my-respect",
			"useHost": false,
			"out":     "iamsalnikov/my-respect",
		},
		{
			"package": "github.com/iamsalnikov/my-respect/subpackage",
			"useHost": false,
			"out":     "iamsalnikov/my-respect",
		},
		{
			"package": "github.com/iamsalnikov/my-respect/subpackage",
			"useHost": true,
			"out":     "github.com/iamsalnikov/my-respect",
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
			"pkgs": []string{
				"os",
				"github.com/iamsalnikov/my-respect",
				"github.com/iamsalnikov/my-respect/subpackage",
				"github.com/some/package/ping",
			},
			"ctrl": map[string]bool{
				"github.com/iamsalnikov/my-respect": true,
				"github.com/some/package":           true,
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

func TestPromptGithubUsernameErr(t *testing.T) {
	_, err := promptGithubUsername(&bytes.Buffer{}, &bytes.Buffer{})
	if err != ErrCanNotGetUsername {
		t.Errorf("I expected to get error \"%s\" but got \"%s\"", ErrCanNotGetUsername, err)
	}
}

func TestPromptGithubUsername(t *testing.T) {
	username := "username"
	in := bytes.NewBufferString(username)
	out := &bytes.Buffer{}

	u, err := promptGithubUsername(out, in)
	if err != nil {
		t.Errorf("I got undexpected error \"%s\"", err)
	}

	if u != username {
		t.Errorf("I expected to get username \"%s\" but got \"%s\"", username, u)
	}
}

func TestPromptGithubToken(t *testing.T) {
	_, err := promptGithubToken(&bytes.Buffer{}, &bytes.Buffer{})
	if err != ErrCanNotGetToken {
		t.Errorf("I expected to get error \"%s\" but got \"%s\"", ErrCanNotGetToken, err)
	}
}

func TestGithubRespecter_promptToken(t *testing.T) {
	token := "token"
	in := bytes.NewBufferString(token)
	out := &bytes.Buffer{}

	tk, err := promptGithubToken(out, in)
	if err != nil {
		t.Errorf("I got undexpected error \"%s\"", err)
	}

	if tk != token {
		t.Errorf("I expected to get username \"%s\" but got \"%s\"", token, tk)
	}
}

func TestGithubRespecter_SayRespectBadResponseStatus(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	githubAPI = server.URL

	github := &GithubRespecter{}

	err := github.SayRespect("")
	if err != ErrCanNotSayRespect {
		t.Errorf("I expected to get error \"%s\" but got \"%s\"", ErrCanNotSayRespect, err)
	}
}

func TestGithubRespecter_SayRespect(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))
	githubAPI = server.URL

	github := &GithubRespecter{}

	err := github.SayRespect("")
	if err != nil {
		t.Errorf("I got unexpected error: %s", err)
	}
}
