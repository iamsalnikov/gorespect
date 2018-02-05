package main

import (
	"bytes"
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

func TestGithubRespecter_promptUsernameErr(t *testing.T) {
	config := NewConfig("")

	github := &GithubRespecter{
		Config: config,
		In:     &bytes.Buffer{},
		Out:    &bytes.Buffer{},
	}

	err := github.promptUsername()
	if err != ErrCanNotGetUsername {
		t.Errorf("I expected to get error \"%s\" but got \"%s\"", ErrCanNotGetUsername, err)
	}
}

func TestGithubRespecter_promptUsername(t *testing.T) {
	config := NewConfig("")

	username := "username"
	in := bytes.NewBufferString(username)
	out := &bytes.Buffer{}

	github := &GithubRespecter{
		Config: config,
		In:     in,
		Out:    out,
	}

	err := github.promptUsername()
	if err != nil {
		t.Errorf("I got undexpected error \"%s\"", err)
	}

	gu, err := config.GetString(githubUserKey)
	if err != nil {
		t.Errorf("I got undexpected error \"%s\"", err)
	}

	if gu != username {
		t.Errorf("I expected to get username \"%s\" but got \"%s\"", username, gu)
	}
}

func TestGithubRespecter_promptTokenErr(t *testing.T) {
	config := NewConfig("")

	github := &GithubRespecter{
		Config: config,
		In:     &bytes.Buffer{},
		Out:    &bytes.Buffer{},
	}

	err := github.promptToken()
	if err != ErrCanNotGetToken {
		t.Errorf("I expected to get error \"%s\" but got \"%s\"", ErrCanNotGetToken, err)
	}
}

func TestGithubRespecter_promptToken(t *testing.T) {
	config := NewConfig("")

	token := "token"
	in := bytes.NewBufferString(token)
	out := &bytes.Buffer{}

	github := &GithubRespecter{
		Config: config,
		In:     in,
		Out:    out,
	}

	err := github.promptToken()
	if err != nil {
		t.Errorf("I got undexpected error \"%s\"", err)
	}

	ct, err := config.GetString(githubTokenKey)
	if err != nil {
		t.Errorf("I got undexpected error \"%s\"", err)
	}

	if ct != token {
		t.Errorf("I expected to get username \"%s\" but got \"%s\"", token, ct)
	}
}

func TestGithubRespecter_SetUpNewUserErr(t *testing.T) {
	config := NewConfig("")

	github := &GithubRespecter{
		Config: config,
		In:     &bytes.Buffer{},
		Out:    &bytes.Buffer{},
	}

	err := github.SetUp()
	if err != ErrCanNotGetUsername {
		t.Errorf("I expected to get error \"%s\" but got \"%s\"", ErrCanNotGetUsername, err)
	}
}

func TestGithubRespecter_SetUpNewUserAndNewTokenErr(t *testing.T) {
	config := NewConfig("")

	username := "username"
	github := &GithubRespecter{
		Config: config,
		In:     bytes.NewBufferString(username),
		Out:    &bytes.Buffer{},
	}

	err := github.SetUp()
	if err != ErrCanNotGetToken {
		t.Errorf("I expected to get error \"%s\" but got \"%s\"", ErrCanNotGetToken, err)
	}

	u, _ := config.GetString(githubUserKey)
	if u != username {
		t.Errorf("I expected to get \"%s\" but got \"%s\"", username, u)
	}
}

func TestGithubRespecter_SetUpNewUserAndNewToken(t *testing.T) {
	config := NewConfig("")

	username := "username"
	token := "token"
	github := &GithubRespecter{
		Config: config,
		In:     bytes.NewBufferString(username + "\n" + token),
		Out:    &bytes.Buffer{},
	}

	err := github.SetUp()
	if err != nil {
		t.Errorf("I got unexpected error: %s", err)
	}

	u, _ := config.GetString(githubUserKey)
	if u != username {
		t.Errorf("I expected to get \"%s\" but got \"%s\"", username, u)
	}

	ct, _ := config.GetString(githubTokenKey)
	if ct != token {
		t.Errorf("I expected to get \"%s\" but got \"%s\"", token, ct)
	}
}

func TestGithubRespecter_SetUpWithExistingData(t *testing.T) {
	config := NewConfig("")
	config.SetValue(githubUserKey, "user")
	config.SetValue(githubTokenKey, "token")

	github := &GithubRespecter{
		Config: config,
		In:     &bytes.Buffer{},
		Out:    &bytes.Buffer{},
	}

	err := github.SetUp()
	if err != nil {
		t.Errorf("I got unexpected error: %s", err)
	}
}
