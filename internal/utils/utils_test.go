package utils

import (
	"fmt"
	"os"
	"regexp"
	"testing"
)

func TestSetEnvFromFile(t *testing.T) {
	SetEnvFromFile()

	env := os.Getenv("ENV")

	if env == "" {
		t.Errorf("env is empty")
	}
}

func TestGetChanges(t *testing.T) {
	out := GetChanges()

	if out == "" {
		t.Error("the output is empty")
	}
}

func TestGetDate(t *testing.T) {
	now, _ := GetDate()
	fmt.Println("message", now)

	reg := regexp.MustCompile(`\d{4}-\d{2}-\d{2}$`)

	if !reg.MatchString(now) {
		t.Error("the date ir wrong", now)
	}
}
