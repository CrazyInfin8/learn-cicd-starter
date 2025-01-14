package main_test

import (
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	_, err := auth.GetAPIKey(nil)
	if err == nil {
		t.Error("expected error, got nil")
	}
}
