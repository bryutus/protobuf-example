package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/bryutus/protobuf-tutorial/tutorial"
	"github.com/golang/protobuf/proto"
)

// Main reads the entire address book from a file, adds one person based on
// user input, then writes it back out to the same file.
func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s ADDRESS_BOOK_FILE\n", os.Args[0])
	}
	fname := os.Args[1]

	// Read the existing address book.
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s: File not found. Creating new file.\n", fname)
		} else {
			log.Fatalln("Error reading file:", err)
		}
	}

	// [START marshal_proto]
	book := &pb.AddressBook{}
	// [START_EXCLUDE]
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
}
