package main

import (
	"os"
	"testing"
)

/*
 * The unit tests in this file simulate command line invocation.
 */

func TestMain(test *testing.T) {
	os.Args = []string{programName, "socket", "--socket-file",  "/tmp/test-mock-server.sock", "--help"}
	main()
}
