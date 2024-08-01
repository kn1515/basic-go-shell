package main

import (
	"fmt"
	"os"
	"strings"
)

var builtinCommandFuncMapping = map[string]func([]string){
    "echo":  builtinEcho,
    "exit":  builtinExit,
    "clear": builtinClear,
}

func builtinEcho(args []string) {
    fmt.Println(strings.Join(args, " "))
}

func builtinExit(args []string) {
    os.Exit(0)
}

func builtinClear(args []string) {
    fmt.Print("\033[H\033[2J\033[3J")
}
