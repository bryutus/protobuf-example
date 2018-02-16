package main

import (
	"strings"
	"testing"
)

func TestPromptAddressReturnsAddress(t *testing.T) {
	in := `12345
Example Name
name@example.com
`
	got, err := promptForAddress(strings.NewReader(in))
	if err != nil {
		t.Fatalf("promptForAddress(%q) had unexpected error: %s", in, err.Error())
	}
	if got.Id != 12345 {
		t.Errorf("promptForAddress(%q) got %d, want ID %d", in, got.Id, 12345)
	}
	if got.Name != "Example Name" {
		t.Errorf("promptForAddress(%q) => want name %q, got %q", in, "Example Name", got.Name)
	}
	if got.Email != "name@example.com" {
		t.Errorf("promptForAddress(%q) => want email %q, got %q", in, "name@example.com", got.Email)
	}
}
