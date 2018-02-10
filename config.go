package main

import (
	"encoding/json"
	"os"
	"sync"
)

// GithubConfig struct represents configuration for GithubRespecter
type GithubConfig struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

// Config struct for working with configuration
type Config struct {
	Github GithubConfig `json:"github"`
}

// ConfigStorage struct saves and loads config
type ConfigStorage struct {
	FilePath string
	m        *sync.RWMutex
}

// NewConfigStorage func creates new config storage
// If we will not pass file path then we work in-memory
func NewConfigStorage(filePath string) *ConfigStorage {
	return &ConfigStorage{
		FilePath: filePath,
		m:        &sync.RWMutex{},
	}
}

// Load function creates new config
func (cs *ConfigStorage) Load() *Config {
	c := &Config{}

	if cs.FilePath == "" {
		return c
	}

	cs.m.RLock()
	defer cs.m.RUnlock()

	file, err := os.Open(cs.FilePath)
	defer file.Close()
	if err != nil {
		return c
	}

	decoder := json.NewDecoder(file)
	decoder.Decode(&c)

	return c
}

// Save func writes config file if we know it's name
func (cs *ConfigStorage) Save(c *Config) error {
	if cs.FilePath == "" {
		return nil
	}

	cs.m.Lock()
	defer cs.m.Unlock()

	file, err := os.Open(cs.FilePath)
	defer file.Close()

	if os.IsNotExist(err) {
		file, err = os.Create(cs.FilePath)
	}

	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(c)
	if err != nil {
		return err
	}

	return nil
}
