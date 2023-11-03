package main

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"os"
)

func main() {
	args := os.Args[1:]
	slog.Info("Starting glox", "args", args)
	if len(args) > 1 {
		slog.Info("Too many arguments", "Usage", "glox [script]")
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}

func runFile(path string) error {
	dat, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return run(string(dat))
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF { // usually sent via ctrl-D
				break
			}
			slog.Error("Error reading line", "err", err)
			continue
		}
		run(line)
	}
}

func run(source string) error {
	// scanner := scanner.NewScanner(source)
	// tokens := scanner.ScanTokens()

	fmt.Println(source)

	// For now, just print the tokens.
	// for _, token := range tokens {
	// }
	return nil
}
