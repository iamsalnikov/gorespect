package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	githubHost     = "github.com"
	githubAPIHost  = "api.github.com"
	githubUserKey  = "github.user"
	githubTokenKey = "github.token"
)

var (
	// ErrCanNotSayRespect shows that we can not give a start go github repo
	ErrCanNotSayRespect = errors.New("can not say respect")
	// ErrCanNotGetUsername shows that we can not get username of github user
	ErrCanNotGetUsername = errors.New("can not get username")
	// ErrCanNotGetToken show that we can not get access token of github user
	ErrCanNotGetToken = errors.New("can not get token")
)

// GithubRespecter works with github packages
type GithubRespecter struct {
	Config *Config
}

// CanProcess func checks if we can work with this package
func (g *GithubRespecter) CanProcess(p string) bool {
	return strings.Index(p, githubHost) == 0
}

// FilterRespectable func returns packages which with we can work
func (g *GithubRespecter) FilterRespectable(pkgs []string) []string {
	pMap := make(map[string]bool)
	for _, p := range pkgs {
		p = g.normalizePackageName(p, true)
		if _, ok := pMap[p]; ok {
			continue
		}

		if g.CanProcess(p) {
			pMap[p] = true
		}
	}

	res := make([]string, len(pMap))
	var i int
	for p := range pMap {
		res[i] = p
		i++
	}

	return res
}

// SetUp func checks config
func (g *GithubRespecter) SetUp() error {
	if !g.Config.HasValue(githubUserKey) {
		err := g.promptUsername()
		if err != nil {
			return err
		}
	}

	if !g.Config.HasValue(githubTokenKey) {
		err := g.promptToken()
		if err != nil {
			return err
		}
	}

	return nil
}

// SayRespect func gives a star to github repos
func (g *GithubRespecter) SayRespect(p string) error {
	client := http.Client{}
	request, err := http.NewRequest(http.MethodPut, "/user/starred/"+g.normalizePackageName(p, false), nil)
	if err != nil {
		return err
	}

	request.URL.Scheme = "https"
	request.URL.Host = githubAPIHost

	user, _ := g.Config.GetString(githubUserKey)
	password, _ := g.Config.GetString(githubTokenKey)

	request.URL.User = url.UserPassword(user, password)

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusNoContent {
		return ErrCanNotSayRespect
	}

	return nil
}

func (g *GithubRespecter) promptUsername() error {
	var username string
	_, err := prompt("Enter github username: ", &username)
	if err != nil {
		return ErrCanNotGetUsername
	}

	g.Config.SetValue(githubUserKey, username)

	return nil
}

func (g *GithubRespecter) promptToken() error {
	var token string
	tokenURL := fmt.Sprintf("https://%s/settings/tokens/new?scopes=public_repo&description=GoMyRespect", githubHost)
	message := fmt.Sprintf("Please, generate and copy token here: %s\nEnter token: ", tokenURL)

	_, err := prompt(message, &token)
	if err != nil {
		return ErrCanNotGetToken
	}

	g.Config.SetValue(githubTokenKey, token)

	return nil
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
