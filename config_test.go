package main

import (
	"os"
	"testing"
)

func TestConfigStorage_LoadFromEmptyFile(t *testing.T) {
	storage := NewConfigStorage("")
	config := storage.Load()

	if len(config.Github.Username) > 0 {
		t.Errorf("I expected to get empty usename")
	}
}

func TestConfigStorage_LoadFromUnexistingFile(t *testing.T) {
	storage := NewConfigStorage("")
	config := storage.Load()

	if len(config.Github.Username) > 0 {
		t.Errorf("I expected to get empty usename")
	}
}

func TestConfigStorage_Load(t *testing.T) {
	storage := NewConfigStorage("testdata/not-empty-config.json")
	config := storage.Load()

	if config.Github.Username != "john" {
		t.Errorf("I expected to get username \"john\" but got \"%s\"", config.Github.Username)
	}
}

func TestConfigStorage_SaveEmptyFileName(t *testing.T) {
	storage := NewConfigStorage("")
	config := &Config{}

	err := storage.Save(config)
	if err != nil {
		t.Errorf("I got unexpected error '%s'", err)
	}
}

func TestConfigStorage_SaveWithNewFileName(t *testing.T) {
	fileName := "testdata/test.json"

	storage := NewConfigStorage(fileName)
	config := &Config{Github: GithubConfig{
		Username: "jack",
		Token:    "daniels",
	}}

	err := storage.Save(config)
	defer os.Remove(fileName)
	if err != nil {
		t.Errorf("I got unexpected error '%s'", err)
	}

	config = storage.Load()
	if config.Github.Username != "jack" {
		t.Errorf("I expected to get username \"jack\" but got \"%s\"", config.Github.Username)
	}
}
