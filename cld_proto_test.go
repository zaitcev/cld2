// Functional test (but being run as a unit test) for CLD protocol support
package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestBasicOpen(t *testing.T) {
	cli_name := "cldcli"
	var cli_path string
	var err error
	cli_path, err = exec.LookPath(cli_name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "No %s found: %s\n", cli_name, err.Error())
		t.Fail()
		return
	}

	cmd := exec.Command(cli_path)
	err = cmd.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Start error: %v\n", err)
		t.Fail()
		return
	}
	// XXX Does it block if command floods its stdout?
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Wait exit error non-nil\n")

		// F23 ships with golang-1.5.3, so this fals with:
		// ./cld_proto_test.go:37: ee.Stderr undefined
		//     (type *exec.ExitError has no field or method Stderr)
		// ee, ok := err.(*exec.ExitError)
		// fmt.Printf("%s\n", string(ee.Stderr))

		// XXX This yields "exit status 1"; the actual error message
		// is not captured as expected. Need to capture error by hand.
		fmt.Printf("%s\n", err.Error())
	} else {
		fmt.Printf("Wait exit error is nil\n")
	}
}
