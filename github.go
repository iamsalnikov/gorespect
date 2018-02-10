package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const githubHost = "github.com"

var (
	// ErrCanNotSayRespect shows that we can not give a start go github repo
	ErrCanNotSayRespect = errors.New("can not say respect")
	// ErrCanNotGetUsername shows that we can not get username of github user
	ErrCanNotGetUsername = errors.New("can not get username")
	// ErrCanNotGetToken show that we can not get access token of github user
	ErrCanNotGetToken = errors.New("can not get token")

	githubAPI = "https://api.github.com"
)

// GithubRespecter works with github packages
type GithubRespecter struct {
	Username string
	Token    string
	Out      io.Writer
	In       io.Reader
}

// CanProcess func checks if we can work with this package
func (g *GithubRespecter) CanProcess(p string) bool {
	return strings.Index(p, githubHost) == 0
}

// FilterRespectable func returns packages which with we can work
func (g *GithubRespecter) FilterRespectable(pkgs []string) []string {
	pMap := make(map[string]bool)
	res := make([]string, 0)

	for _, p := range pkgs {
		p = g.normalizePackageName(p, true)
		if pMap[p] {
			continue
		}

		if g.CanProcess(p) {
			pMap[p] = true
			res = append(res, p)
		}
	}

	return res
}

// SayRespect func gives a star to github repos
func (g *GithubRespecter) SayRespect(p string) error {
	client := http.Client{}
	request, err := http.NewRequest(http.MethodPut, githubAPI+"/user/starred/"+g.normalizePackageName(p, false), nil)
	if err != nil {
		return err
	}

	request.URL.User = url.UserPassword(g.Username, g.Token)

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusNoContent {
		return ErrCanNotSayRespect
	}

	return nil
}

func promptGithubUsername(out io.Writer, in io.Reader) (string, error) {
	var username string
	_, err := prompt("Enter github username: ", &username, out, in)
	if err != nil {
		return "", ErrCanNotGetUsername
	}

	return username, nil
}

func promptGithubToken(out io.Writer, in io.Reader) (string, error) {
	var token string
	tokenURL := fmt.Sprintf("https://%s/settings/tokens/new?scopes=public_repo&description=GoRespect", githubHost)
	message := fmt.Sprintf("Please, generate and copy token here: %s\nEnter token: ", tokenURL)

	_, err := prompt(message, &token, out, in)
	if err != nil {
		return "", ErrCanNotGetToken
	}

	return token, nil
}

func (g *GithubRespecter) normalizePackageName(p string, useHost bool) string {
	parts := strings.Split(p, "/")
	if len(parts) < 3 {
		return p
	}

	start := 1
	if useHost {
		start = 0
	}

	return strings.Join(parts[start:3], "/")
}
