package helloworld

import (
	"github.com/docktermj/go-hello-world-plus/common/runner"
	"github.com/docktermj/go-hello-world-plus/subcommand/helloworld/english"
	"github.com/docktermj/go-hello-world-plus/subcommand/helloworld/german"
)

func Command(argv []string) {

	usage := `
Usage:
    go-hello-world-plus hello-world <subcommand> [<args>...]

Subcommands:
    english    Hello, World!
    german     Hallo, Welt!
`

	functions := map[string]interface{}{
		"english": english.Command,
		"german":  german.Command,
	}

	runner.Run(argv, functions, usage)
}
