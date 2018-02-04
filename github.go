package main

import (
	"strings"
	"net/url"
	"net/http"
	"fmt"
	"errors"
)

const (
	GithubHost = "github.com"
	GithubApiHost = "api.github.com"
	GithubUserKey = "github.user"
	GithubTokenKey = "github.token"
)

var (
	ErrCanNotSayRespect = errors.New("can not say respect")
	ErrCanNotGetUsername = errors.New("can not get username")
	ErrCanNotGetToken = errors.New("can not get token")
)

type GithubRespecter struct {
	Config *Config
}

func (g *GithubRespecter) CanProcess(p string) bool {
	return strings.Index(p, GithubHost) == 0
}

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

func (g *GithubRespecter) SetUp() error {
	if !g.Config.HasValue(GithubUserKey) {
		err := g.promptUsername()
		if err != nil {
			return err
		}
	}

	if !g.Config.HasValue(GithubTokenKey) {
		err := g.promptToken()
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *GithubRespecter) SayRespect(p string) error {
	client := http.Client{}
	request, err := http.NewRequest(http.MethodPut, "/user/starred/" + g.normalizePackageName(p, false), nil)
	if err != nil {
		return err
	}

	request.URL.Scheme = "https"
	request.URL.Host = GithubApiHost

	user, _ := g.Config.GetString(GithubUserKey)
	password, _ := g.Config.GetString(GithubTokenKey)

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

	g.Config.SetValue(GithubUserKey, username)

	return nil
}

func (g *GithubRespecter) promptToken() error {
	var token string
	tokenUrl := fmt.Sprintf("https://%s/settings/tokens/new?scopes=public_repo&description=GoMyRespect", GithubHost)
	message := fmt.Sprintf("Please, generate and copy token here: %s\nEnter token: ", tokenUrl)

	_, err := prompt(message, &token)
	if err != nil {
		return ErrCanNotGetToken
	}

	g.Config.SetValue(GithubTokenKey, token)

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

