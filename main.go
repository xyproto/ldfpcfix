// A program that takes ld as the first argument and silences the annoying FPC warning

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	wd = []byte("%P: warning: %s contains output sections")
)

func main() {
	in_filename := "/usr/bin/ld"
	out_filename := "/usr/bin/ld"

	// Read the input filename
	data, err := ioutil.ReadFile(in_filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Check if ld has already been patched
	if !bytes.Contains(data, wd) {
		fmt.Println(in_filename, "has already been patched")
		os.Exit(1)
		return
	}
	// Find the position of the warning
	pos := bytes.Index(data, wd)

	// Patch it
	data[pos] = 0 // Silence the message with a 0 byte

	// Get the permissions of the original file
	fi, err := os.Stat(in_filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	perm := fi.Mode().Perm()

	fmt.Printf("Patching %s... ", out_filename)

	// Write the patched data to the new file, but with the same permissions as the original file
	err = ioutil.WriteFile(out_filename, data, perm)
	if err != nil {
		fmt.Println("fail")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("ok")
}
