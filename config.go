package main

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

var ValueNotFound = errors.New("value not found")

type Config struct {
	filePath  string
	config    map[string]interface{}
	dataMutex *sync.Mutex
}

func NewConfig(filePath string) *Config {
	c := &Config{
		filePath:  filePath,
		dataMutex: &sync.Mutex{},
		config:    make(map[string]interface{}),
	}

	if filePath == "" {
		return c
	}

	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return c
	}

	decoder := json.NewDecoder(file)
	decoder.Decode(&c.config)

	return c
}

func (c *Config) HasValue(key string) bool {
	if c.config == nil {
		return false
	}

	_, ok := c.config[key]

	return ok
}

func (c *Config) SetValue(key string, value interface{}) {
	c.dataMutex.Lock()
	c.config[key] = value
	c.dataMutex.Unlock()
}

func (c *Config) GetValue(key string) (interface{}, error) {
	if c.HasValue(key) {
		return c.config[key], nil
	}

	return nil, ValueNotFound
}

func (c *Config) GetString(key string) (string, error) {
	value, err := c.GetValue(key)
	str := ""

	if value != nil {
		str = value.(string)
	}

	return str, err
}

func (c *Config) Save() error {
	if c.filePath == "" {
		return nil
	}

	c.dataMutex.Lock()
	defer c.dataMutex.Unlock()

	file, err := os.Open(c.filePath)
	defer file.Close()

	if os.IsNotExist(err) {
		file, err = os.Create(c.filePath)
	}

	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(&c.config)
	if err != nil {
		return err
	}

	return nil
}
