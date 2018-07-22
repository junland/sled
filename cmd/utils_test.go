package cmd

import (
	"os"
	"testing"
)

func TestGetEnvString(t *testing.T) {
	os.Setenv("TEST_STRING", "this-is-a-string")
	value := GetEnvString("TEST_STRING", "backup")
	if value != "this-is-a-string" {
		t.Errorf("environment variable value is incorrect, got %s", value)
	}

	os.Setenv("TEST_STRING", "")
	value = GetEnvString("TEST_STRING", "backup")
	if value != "backup" {
		t.Errorf("environment variable backup value is incorrect, got %s", value)
	}
}

func TestGetEnvBool(t *testing.T) {
	os.Setenv("TEST_BOOL", "true")
	value := GetEnvBool("TEST_BOOL", false)
	if value != true {
		t.Errorf("environment variable value is incorrect, got %t", value)
	}

	os.Setenv("TEST_BOOL", "")
	value = GetEnvBool("TEST_BOOL", true)
	if value != true {
		t.Errorf("environment variable backup value is incorrect, got %t", value)
	}
}
