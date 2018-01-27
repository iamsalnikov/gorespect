package main

import "strings"

const GithubHost = "github.com"

type Thanker interface {
	CanProcess(p string) bool
	SayThankYou(p string) error
}

type GithubThanker struct {}

func (g *GithubThanker) CanProcess(p string) bool {
	return strings.Index(p, GithubHost) == 0
}

func (g *GithubThanker) SayThankYou(p string) error {
	panic("implement me")
}


