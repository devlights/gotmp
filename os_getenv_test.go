package main

import (
	"testing"
	"os"
)

func TestOsGetenvExists(t *testing.T) {
	// Arrange
	envkey := "GOPATH"

	// Act
	envval := os.Getenv(envkey)

	// Assert
	if envval == "" {
		t.Errorf("存在しない")
	}
}

func TestOsGetenvNotExists(t *testing.T) {
	// Arrange
	envkey := "NOT_EXISTS"

	// Act
	envval := os.Getenv(envkey)

	// Assert
	if envval != "" {
		t.Errorf("存在する")
	}
}