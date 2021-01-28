// ldfpcfix silences the annoying warning in /usr/bin/ld
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

// patchAway takes a filename and a string
// If the string is found in the file, the first byte is
// set to 0, to make the string zero length in C.
func patchAway(filename, cstring string) error {
	// Read the input filename
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// Find the position of the warning
	pos := bytes.Index(data, []byte(cstring))

	// If it does not exist, the file has most likely already been patched
	if pos == -1 {
		return fmt.Errorf("%s has already been patched", filename)
	}

	// Silence the message with a 0 byte
	data[pos] = 0

	// Retrieve the permissions of the original file
	fi, err := os.Stat(filename)
	if err != nil {
		return err
	}
	perm := fi.Mode().Perm()

	// Write the patched data to the new file, but with the same permissions as the original file
	return ioutil.WriteFile(filename, data, perm)
}

func main() {
	const (
		filename       = "/usr/bin/ld"
		warningMessage = "%P: warning: %s contains output sections"
	)
	fmt.Printf("Patching %s... ", filename)
	if err := patchAway(filename, warningMessage); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("ok")
}
