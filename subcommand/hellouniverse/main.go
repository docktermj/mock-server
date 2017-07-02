package hellouniverse

import (
	"github.com/docktermj/go-hello-world-plus/common/runner"
	"github.com/docktermj/go-hello-world-plus/subcommand/hellouniverse/english"
	"github.com/docktermj/go-hello-world-plus/subcommand/hellouniverse/german"
)

func Command(argv []string) {

	usage := `
Usage:
    go-hello-world-plus hello-universe <subcommand> [<args>...]

Subcommands:
    english    Hello, Universe!
    german     Hallo, Universen!
`

	functions := map[string]interface{}{
		"english": english.Command,
		"german":  german.Command,
	}

	runner.Run(argv, functions, usage)
}
