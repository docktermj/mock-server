package socket

// Inspirations:
//  - https://gist.github.com/hakobe/6f70d69b8c5243117787fd488ae7fbf2

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/docktermj/mock-server/common/help"
	"github.com/docopt/docopt-go"
)

// Read a message from the network and respond.
func echoServer(networkConnection net.Conn) {
	for {
		byteBuffer := make([]byte, 512)

		// Read the Unix Domain Socket.

		numberOfBytesRead, err := networkConnection.Read(byteBuffer)
		if err != nil {
			return
		}
		data := byteBuffer[0:numberOfBytesRead]

		// Print what was received over the socket.

		println("Server got:", string(data))

		// Write a response to the Unix Domain Socket.

		_, err = networkConnection.Write(data)
		if err != nil {
			log.Fatal("Writing client error: ", err)
		}
	}
}

// Function for the "command pattern".
func Command(argv []string) {

	usage := `
Usage:
    mock-server socket [options] 

Options:
   -h, --help
   --socket-file=<file>        Socket file
   --debug               Log debugging messages
`

	// DocOpt processing.

	args, _ := docopt.Parse(usage, nil, true, "", false)

	// Test for required commandline options.

	if args["--socket-file"] == nil {
		message := "Missing '--socket-file' parameter"
		fmt.Println(message)
		help.ShowHelp(usage)
		log.Fatalln(message)
	}

	// Get commandline options.

	socketFile := args["--socket-file"].(string)
	isDebug := args["--debug"].(bool)

	// Listen on the Unix Domain Socket.

	if isDebug {
		log.Printf("Starting echo server on %s", socketFile)
	}

	listener, err := net.Listen("unix", socketFile)
	if err != nil {
		log.Fatal("Listen error: ", err)
	}

	// Configure listener to exit when program ends.

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func(listener net.Listener, c chan os.Signal) {
		sig := <-c
		log.Printf("Caught signal %s: shutting down.", sig)
		listener.Close()
		os.Exit(0)
	}(listener, sigc)

	// Read and Echo loop.

	for {
		networkConnection, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}
		go echoServer(networkConnection)
	}
}
