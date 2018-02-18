package main

import (
	"strings"
	"testing"

	pb "github.com/bryutus/protobuf-tutorial/tutorial"
)

func TestPromptAddressReturnsAddress(t *testing.T) {
	in := `12345
Example Name
name@example.com
123-456-7890
222-222-2222
111-111-1111
777-777-7777

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

	want := []*pb.Person_PhoneNumber{
		{Number: "123-456-7890"},
		{Number: "222-222-2222"},
		{Number: "111-111-1111"},
		{Number: "777-777-7777"},
	}

	if len(got.Phones) != len(want) {
		t.Errorf("want %d phone numbers, got %d", len(want), len(got.Phones))
	}
}
