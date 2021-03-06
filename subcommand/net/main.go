package net

// Inspirations:
//  - https://gist.github.com/hakobe/6f70d69b8c5243117787fd488ae7fbf2

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/docktermj/mock-server/common/help"
	"github.com/docopt/docopt-go"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// Return a pseudo-random string of given length.
func randStringRunes(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// Append a pseudo-unique character string to the prefix.
func getServerId() string {
	return randStringRunes(6)
}

// Read a message from the network and respond.
func echoServer(serverId string, networkConnection net.Conn) {
	for {
		byteBuffer := make([]byte, 512)

		// Read from network connection.

		numberOfBytesRead, err := networkConnection.Read(byteBuffer)
		if err != nil {
			return
		}
		inboundMessage := byteBuffer[0:numberOfBytesRead]

		// Print messages.

		fmt.Printf("Receive: %s\n", string(inboundMessage))
		outboundMessage := fmt.Sprintf("Server-%s responds by echoing: \"%s\"", serverId, inboundMessage)
		fmt.Printf("Respond: %s\n\n", outboundMessage)

		// Write to network connection.

		_, err = networkConnection.Write([]byte(outboundMessage))
		if err != nil {
			log.Fatal("Writing client error: ", err)
		}
	}
}

// Function for the "command pattern".
func Command(argv []string) {

	usage := `
Usage:
    mock-server net [options]

Options:
   -h, --help
   --network=<network_type>  Type of network used for communication
   --address=<address>       Address for network_type. Example (unix) /tmp/tmp.sock
   --debug                   Log debugging messages

Where:
   network_type   Examples: 'unix', 'tcp'
   address        Examples: '/tmp/test.sock', '127.0.0.1:12345'
`

	// DocOpt processing.

	args, _ := docopt.Parse(usage, nil, true, "", false)

	// Test for required commandline options.

	message := ""

	if args["--network"] == nil {
		message += "Missing '--network' parameter;"
	}

	if args["--address"] == nil {
		message += "Missing '--address' parameter;"
	}

	if len(message) > 0 {
		help.ShowHelp(usage)
		fmt.Println(strings.Replace(message, ";", "\n", -1))
		log.Fatalln(strings.Replace(message, ";", "; ", -1))
	}

	// Get commandline options.

	network := args["--network"].(string)
	address := args["--address"].(string)
	isDebug := args["--debug"].(bool)

	// Debugging information.

	if isDebug {
		log.Printf("Starting echo server on '%s' network with address '%s'", network, address)
	}

	// Listen on the network connection.

	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatal("Listen error: ", err)
	}

	// Configure listener to exit when program ends.

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func(listener net.Listener, c chan os.Signal) {
		sig := <-c
		log.Printf("Caught signal %s: shutting down.", sig)
		listener.Close()
		os.Exit(0)
	}(listener, sigc)

	// Read and Echo loop.

	serverId := getServerId()
	for {
		networkConnection, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}
		go echoServer(serverId, networkConnection)
	}
}
