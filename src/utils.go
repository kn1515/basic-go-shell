package main

import (
	"fmt"
	"os"
	"strings"
)

func interpretLine(line string) {
    if strings.HasPrefix(strings.TrimSpace(line), "#") {
        return
    }

    tokens := strings.Fields(line)

    if len(tokens) == 0 {
        return
    }

    redirectionCount := strings.Count(line, ">")
    pipeCount := strings.Count(line, "|")

    if redirectionCount+pipeCount >= 2 {
        fmt.Fprintln(os.Stderr, "More than 2 redirections or pipes are not supported")
        return
    }

    if redirectionCount == 1 {
        redirectionIndex := strings.Index(line, ">")
        commandAndArgs := tokens[:redirectionIndex]
        stdoutFileName := tokens[redirectionIndex+1]

        stdoutFile, err := os.Create(stdoutFileName)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
            return
        }
        defer stdoutFile.Close()

        runCommand(commandAndArgs, nil, stdoutFile)

    } else if pipeCount == 1 {
        pipeIndex := strings.Index(line, "|")
        commandAndArgs1 := tokens[:pipeIndex]
        commandAndArgs2 := tokens[pipeIndex+1:]

        pipeR, pipeW, err := os.Pipe()
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error creating pipe: %v\n", err)
            return
        }
        defer pipeR.Close()
        defer pipeW.Close()

        runCommand(commandAndArgs1, nil, pipeW)
        runCommand(commandAndArgs2, pipeR, nil)

    } else {
        runCommand(tokens, nil, nil)
    }
}
