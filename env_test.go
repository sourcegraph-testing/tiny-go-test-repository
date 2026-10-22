package main

import (
	"os"
	"testing"
)

func TestDogfoodEnv(t *testing.T) {
	if got := os.Getenv("DOGFOOD_35"); got != "1" {
		t.Fatalf("expected DOGFOOD_35=1, got %q", got)
	}
}
