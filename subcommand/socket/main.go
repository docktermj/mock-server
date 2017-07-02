package socket

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/docopt/docopt-go"
)

func echoServer(c net.Conn) {
	for {
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:nr]
		println("Server got:", string(data))
		_, err = c.Write(data)
		if err != nil {
			log.Fatal("Writing client error: ", err)
		}
	}
}

// Function for the "command pattern".
func Command(argv []string) {

	usage := `
Usage:
    domain-socket-tester socket [options] 

Options:
   -h, --help
   --socket-file=<file>        Socket file
   --debug               Log debugging messages
`

	// DocOpt processing.

	args, _ := docopt.Parse(usage, nil, true, "", false)
	socketFile := args["--socket-file"]
	isDebug := args["--debug"]

	// Listen on the Unix Domain Socket

	if isDebug {
		log.Printf("Starting echo server on ", socketFile)
	}

	ln, err := net.Listen("unix", socketFile)
	if err != nil {
		log.Fatal("Listen error: ", err)
	}

	//

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)
	go func(ln net.Listener, c chan os.Signal) {
		sig := <-c
		log.Printf("Caught signal %s: shutting down.", sig)
		ln.Close()
		os.Exit(0)
	}(ln, sigc)

	for {
		fd, err := ln.Accept()
		if err != nil {
			log.Fatal("Accept error: ", err)
		}

		go echoServer(fd)
	}

}
