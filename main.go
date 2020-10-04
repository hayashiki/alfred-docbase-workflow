package main

import (
	"fmt"
	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
	"github.com/hayashiki/alfred-docbase-workflow/handler"
	"log"
	"os"
	"unicode"
)

const repo = "hayashiki/alfred-docbase-workflow"

func main() {
	wf := aw.New(update.GitHub(repo))
	wf.Run(func() {
		run(wf)
	})
}

func run(wf *aw.Workflow) {
	args := wf.Args()
	if len(args) == 0 {
		exitWithError("please provide some input 👀")
	}

	handlers := map[string]func(*aw.Workflow, string) (string, error){
		"setup":   handler.DoSetup,
		"cleanup": handler.DoCleanup,
		"update":  handler.DoUpdate,
		"install": handler.DoInstall,
		"search":  handler.DoSearch,
	}

	cmd, arg := parseArgs(args)

	h, found := handlers[cmd]

	if !found {
		exitWithError(cmd)
	}

	msg, err := h(wf, arg)
	if err != nil {
		exitWithError(err.Error())
		os.Exit(1)
	}

	if msg != "" {
		fmt.Print(msg)
	}
}

func exitWithError(msg string) {
	fmt.Print(capitalize(msg))
	log.Print(msg)
	os.Exit(1)
}

func capitalize(msg string) string {
	return string(unicode.ToUpper(rune(msg[0]))) + msg[1:]
}

func parseArgs(args []string) (cmd, arg string) {
	if len(args) == 1 {
		return args[0], ""
	}

	if len(args) == 2 {
		return args[0], args[1]
	}

	return "", ""
}
