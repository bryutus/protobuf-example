package main

import (
	"bytes"
	"testing"

	pb "github.com/bryutus/protobuf-tutorial/tutorial"
)

func TestWritePersonWritesPerson(t *testing.T) {
	buf := new(bytes.Buffer)
	// [START populate_proto]
	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}
	// [END populate_proto]

	writePerson(buf, &p)
	got := buf.String()
	want := `Person ID: 1234
  Name: John Doe
  E-mail address: jdoe@example.com
  Home Phone #: 555-4321
`

	if got != want {
		t.Errorf("writePerson(%s) =>\n\t%q, want %q", p.String(), got, want)
	}
}
