// ldfpcfix silences the annoying warning in /usr/bin/ld
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	inFilename  = "/usr/bin/ld"
	outFilename = "/usr/bin/ld"
)

var (
	warningMessage = []byte("%P: warning: %s contains output sections")
)

func main() {
	// Read the input filename
	data, err := ioutil.ReadFile(inFilename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Check if ld has already been patched
	if !bytes.Contains(data, warningMessage) {
		fmt.Println(inFilename, "has already been patched")
		os.Exit(1)
		return
	}

	// Find the position of the warning
	pos := bytes.Index(data, warningMessage)

	// Patch it
	data[pos] = 0 // Silence the message with a 0 byte

	// Get the permissions of the original file
	fi, err := os.Stat(inFilename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	perm := fi.Mode().Perm()

	// Write the patched data to the new file, but with the same permissions as the original file
	fmt.Printf("Patching %s... ", outFilename)
	err = ioutil.WriteFile(outFilename, data, perm)
	if err != nil {
		fmt.Println("fail")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("ok")
}
