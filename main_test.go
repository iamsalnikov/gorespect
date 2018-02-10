package main

import (
	"bytes"
	"testing"
)

func TestSetUpConfigEmptyUsernameErr(t *testing.T) {
	config := &Config{}

	in := &bytes.Buffer{}
	out := &bytes.Buffer{}

	err := setUpConfig(config, out, in)
	if err == nil {
		t.Errorf("I expected to get some error, but got nothing")
	}
}

func TestSetUpConfigEmptyUsernameTokenErr(t *testing.T) {
	config := &Config{}

	in := bytes.NewBufferString("username")
	out := &bytes.Buffer{}

	err := setUpConfig(config, out, in)
	if err == nil {
		t.Errorf("I expected to get some error, but got nothing")
	}
}

func TestSetUpConfigEmptyTokenErr(t *testing.T) {
	config := &Config{
		Github: GithubConfig{
			Username: "hello",
		},
	}

	in := &bytes.Buffer{}
	out := &bytes.Buffer{}

	err := setUpConfig(config, out, in)
	if err == nil {
		t.Errorf("I expected to get some error, but got nothing")
	}
}

func TestSetUpConfigEmptyToken(t *testing.T) {
	config := &Config{
		Github: GithubConfig{
			Username: "hello",
		},
	}

	in := bytes.NewBufferString("token")
	out := &bytes.Buffer{}

	err := setUpConfig(config, out, in)
	if err != nil {
		t.Errorf("I got unexpected error: %s", err)
	}
}

func TestSetUpConfigNewData(t *testing.T) {
	config := &Config{}

	in := bytes.NewBufferString("username\ntoken")
	out := &bytes.Buffer{}

	err := setUpConfig(config, out, in)
	if err != nil {
		t.Errorf("I got unexpected error: %s", err)
	}

	if config.Github.Username != "username" {
		t.Errorf("I expected to get \"username\" but got \"%s\"", config.Github.Username)
	}

	if config.Github.Token != "token" {
		t.Errorf("I expected to get \"token\" but got \"%s\"", config.Github.Token)
	}
}

func TestSetUpExistingNewData(t *testing.T) {
	config := &Config{
		Github: GithubConfig{
			Username: "username",
			Token:    "token",
		},
	}

	in := bytes.NewBufferString("user\ntok")
	out := &bytes.Buffer{}

	err := setUpConfig(config, out, in)
	if err != nil {
		t.Errorf("I got unexpected error: %s", err)
	}

	if config.Github.Username != "username" {
		t.Errorf("I expected to get \"username\" but got \"%s\"", config.Github.Username)
	}

	if config.Github.Token != "token" {
		t.Errorf("I expected to get \"token\" but got \"%s\"", config.Github.Token)
	}
}
