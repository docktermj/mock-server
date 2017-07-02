package socket

import (
    "fmt"
    
	"github.com/docopt/docopt-go"
)

// Function for the "command pattern".
func Command(argv []string) {

	usage := `
Usage:
    domain-socket-tester socket [options] 

Options:
   -h, --help
   --socket-file=<file>        Socket file
`

	// DocOpt processing.

	args, _ := docopt.Parse(usage, nil, true, "", false)
	
	fmt.Println(args)




}
