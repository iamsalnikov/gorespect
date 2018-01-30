package main

import (
	"strings"
	"net/url"
	"net/http"
	"fmt"
	"os"
)

const GithubHost = "github.com"
const GithubApiHost = "api.github.com"
const GithubUserKey = "github.user"
const GithubTokenKey = "github.token"

type GithubThanker struct {
	Config *Config
}

func (g *GithubThanker) CanProcess(p string) bool {
	return strings.Index(p, GithubHost) == 0
}

func (g *GithubThanker) SayThankYou(p string) error {
	if !g.Config.HasValue(GithubUserKey) {
		g.promptUsername()
	}

	if !g.Config.HasValue(GithubTokenKey) {
		g.promptToken()
	}

	client := http.Client{}
	request, err := http.NewRequest(http.MethodPut, "/user/starred/" + g.normalizePackageName(p), nil)
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
		return ErrCanNotSayThanks
	}

	return nil
}

func (g *GithubThanker) promptUsername() {
	var username string
	_, err := Prompt("Enter github username: ", &username)
	if err != nil {
		fmt.Println("Can not get username")
		os.Exit(1)
	}

	g.Config.SetValue(GithubUserKey, username)
}

func (g *GithubThanker) promptToken() {
	var token string
	tokenUrl := fmt.Sprintf("https://%s/settings/tokens/new?scopes=public_repo&description=GoThanks", GithubHost)
	message := fmt.Sprintf("Please, generate and copy token here: %s\nEnter token: ", tokenUrl)

	_, err := Prompt(message, &token)
	if err != nil {
		fmt.Println("Can not get token")
		os.Exit(1)
	}

	g.Config.SetValue(GithubTokenKey, token)
}

func (g *GithubThanker) normalizePackageName(p string) string {
	parts := strings.Split(p, "/")
	if len(parts) < 3 {
		return p
	}

	return strings.Join(parts[1:3], "/")
}

