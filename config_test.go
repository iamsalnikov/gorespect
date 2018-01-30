package main

import (
	"testing"
	"os"
)

func TestNewConfigWithEmptyFileName(t *testing.T) {
	config := NewConfig("")

	if len(config.config) > 0 {
		t.Errorf("I expected to get empty config, but got %+v", config.config)
	}
}

func TestNewConfigWithUnexistingFile(t *testing.T) {
	config := NewConfig("testdata/unexist.json")

	if len(config.config) > 0 {
		t.Errorf("I expected to get empty config, but got %+v", config.config)
	}
}

func TestNewConfigWithExistingFile(t *testing.T) {
	config := NewConfig("testdata/not-empty-config.json")

	name, err := config.GetString("name")
	if err != nil {
		t.Errorf("I got unexpected error '%s'", err)
	}

	if name != "John" {
		t.Errorf("I expected to get name 'John' but got '%s'", name)
	}
}

func TestConfig_SetValue(t *testing.T) {
	config := NewConfig("")
	config.SetValue("test", "me")

	actual := config.config["test"].(string)

	if actual != "me" {
		t.Errorf("I expected to get \"me\" but got \"%s\"", actual)
	}
}

func TestConfig_HasValueNewConfig(t *testing.T) {
	config := NewConfig("")
	has := config.HasValue("test")

	if has {
		t.Errorf("I expected to get false but got true")
	}
}

func TestConfig_HasValueUnexisting(t *testing.T) {
	config := NewConfig("")
	config.SetValue("me", "test")
	has := config.HasValue("test")

	if has {
		t.Errorf("I expected to get false but got true")
	}
}

func TestConfig_HasValueExisting(t *testing.T) {
	config := NewConfig("")
	config.SetValue("me", "test")
	has := config.HasValue("me")

	if !has {
		t.Errorf("I expected to get true but got false")
	}
}

func TestConfig_GetValueNewConfig(t *testing.T) {
	config := NewConfig("")
	val, err := config.GetValue("test")

	if err != ValueNotFound {
		t.Errorf("I expected to get error \"%s\" but got \"%v\"", ValueNotFound, err)
	}

	if val != nil {
		t.Errorf("I expected to get value %v but got %v", nil, val)
	}
}

func TestConfig_GetValueUnexisting(t *testing.T) {
	config := NewConfig("")
	config.SetValue("me", "test")
	val, err := config.GetValue("test")

	if err != ValueNotFound {
		t.Errorf("I expected to get error \"%s\" but got \"%v\"", ValueNotFound, err)
	}

	if val != nil {
		t.Errorf("I expected to get value %v but got %v", nil, val)
	}
}

func TestConfig_GetValueExisting(t *testing.T) {
	config := NewConfig("")
	config.SetValue("test", "me")
	val, err := config.GetValue("test")

	if err != nil {
		t.Errorf("I got unexpected error %s", err)
	}

	if val.(string) != "me" {
		t.Errorf("I expected to get value 'me' but got '%v'", val.(string))
	}
}

func TestConfig_GetStringNewConfig(t *testing.T) {
	config := NewConfig("")
	val, err := config.GetString("test")

	if err != ValueNotFound {
		t.Errorf("I expected to get error \"%s\" but got \"%s\"", ValueNotFound, err)
	}

	if val != "" {
		t.Errorf("I expected to get value '%s' but got '%s'", nil, val)
	}
}

func TestConfig_GetStringUnexisting(t *testing.T) {
	config := NewConfig("")
	config.SetValue("me", "test")
	val, err := config.GetString("test")

	if err != ValueNotFound {
		t.Errorf("I expected to get error \"%s\" but got \"%v\"", ValueNotFound, err)
	}

	if val != "" {
		t.Errorf("I expected to get value '%s' but got '%s'", nil, val)
	}
}

func TestConfig_GetStringExisting(t *testing.T) {
	config := NewConfig("")
	config.SetValue("test", "me")
	val, err := config.GetString("test")

	if err != nil {
		t.Errorf("I got unexpected error %s", err)
	}

	if val != "me" {
		t.Errorf("I expected to get value 'me' but got '%s'", val)
	}
}

func TestConfig_SaveWithEmptyFileName(t *testing.T) {
	config := NewConfig("")
	config.SetValue("test", "me")

	err := config.Save()
	if err != nil {
		t.Errorf("I got unexpected error '%s'", err)
	}
}

func TestConfig_SaveWithNewFileName(t *testing.T) {
	fileName := "testdata/test.json"
	config := NewConfig(fileName)

	config.SetValue("test", "me")
	err := config.Save()
	defer os.Remove(fileName)
	if err != nil {
		t.Errorf("I got unexpected error '%s'", err)
	}

	config = NewConfig(fileName)
	val, err := config.GetString("test")
	if err != nil {
		t.Errorf("I got unexpected error '%s'", err)
	}

	if val != "me" {
		t.Errorf("I expected value 'me' but got '%s'", val)
	}
}