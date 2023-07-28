package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
import log "github.com/inconshreveable/log15"

func main() {
	argsWithoutProg := os.Args[1:]
	log.Info("Starting glox", "args", argsWithoutProg)
	if len(argsWithoutProg) > 1 {
		log.Info("Too many arguments")
		os.Exit(64)
	} else if len(argsWithoutProg) == 1 {
		log.Info("Running File", "path", argsWithoutProg[0])
		e := runFile(argsWithoutProg[0])
		if e != nil {
			os.Exit(65)
		}
		os.Exit(0)
	} else {
		log.Info("Running in REPL mode")
		e := runPrompt()
		if e != nil {
			os.Exit(65)
		}
		os.Exit(0)
	}
}

func runFile(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return run(string(bytes))
}

func runPrompt() error {
	for {
		fmt.Print("> : ")
		var input string
		if _, e := fmt.Scanln(&input); e != nil {
			log.Warn("Input scan error", "err", e)
		}
		e := run(input)
		if e != nil {
			log.Warn("Error running input", "err", e, "input", input)
		}
	}
}

func run(source string) error {
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		log.Info(scanner.Text())
	}
	if e := scanner.Err(); e != nil {
		reportError(0, "", e.Error())
		return e
	}
	return nil
}

// TODO: this should be some kind of sentinel error / ErrorReporter interface
func reportError(line uint, where string, message string) {
	log.Warn("Execution was not a gooder", "line", line, "location", where, "message", message)
}
