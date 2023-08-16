package main

import (
	"bufio"
	"glox/pkg/lex"
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
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(lex.Tokenizer)
	log.Info("Please enter text. Enter 'exit' to quit.")
	for {
		for scanner.Scan() {
			if scanner.Text() == "exit" {
				break
			}
			log.Info(scanner.Text())
		}
		if e := scanner.Err(); e != nil {
			reportError(0, "", e.Error())
			continue // TODO: if running in REPL mode, mistake in source code should not kill the entire session
		}
		return nil
	}
}

func run(source string) error {
	scanner := bufio.NewScanner(strings.NewReader(source))
	scanner.Split(lex.Tokenizer)
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
	log.Warn("Execution failed", "line", line, "location", where, "message", message)
}
