package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func findCommand(command string) string {
    if strings.Contains(command, "/") {
        if _, err := os.Stat(command); err == nil {
            return command
        }
        return ""
    }

    pathDirs := strings.Split(os.Getenv("PATH"), ":")
    for _, pathDir := range pathDirs {
        commandPath := pathDir + "/" + command
        if stat, err := os.Stat(commandPath); err == nil && !stat.IsDir() {
            return commandPath
        }
    }
    return ""
}

func setStdio(stdinFd, stdoutFd *os.File) {
    if stdinFd != nil {
        syscall.Dup2(int(stdinFd.Fd()), int(os.Stdin.Fd()))
    }
    if stdoutFd != nil {
        syscall.Dup2(int(stdoutFd.Fd()), int(os.Stdout.Fd()))
    }
}

func runCommand(tokens []string, stdinFd, stdoutFd *os.File) {
    command := tokens[0]
    args := tokens[1:]

    if fn, ok := builtinCommandFuncMapping[command]; ok {
        defaultStdinDupFd := os.Stdin
        defaultStdoutDupFd := os.Stdout
        setStdio(stdinFd, stdoutFd)

        fn(args)

        setStdio(defaultStdinDupFd, defaultStdoutDupFd)
    } else {
        commandPath := findCommand(command)

        if commandPath == "" {
            fmt.Fprintf(os.Stderr, "%s: command not found\n", command)
            return
        }

        cmd := exec.Command(commandPath, args...)
        if stdinFd != nil {
            cmd.Stdin = stdinFd
        }
        if stdoutFd != nil {
            cmd.Stdout = stdoutFd
        }
        cmd.Stderr = os.Stderr

        if err := cmd.Run(); err != nil {
            fmt.Fprintf(os.Stderr, "Error running command: %v\n", err)
        }
    }
}
