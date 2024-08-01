package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    if len(os.Args) >= 2 {
        fileName := os.Args[1]
        file, err := os.Open(fileName)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
            return
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            interpretLine(scanner.Text())
        }
        if err := scanner.Err(); err != nil {
            fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
        }

    } else {
        reader := bufio.NewReader(os.Stdin)
        for {
            fmt.Print("🐚 \033[33mbasic-go-sh\033[0m > ")
            line, _ := reader.ReadString('\n')
            interpretLine(line)
        }
    }
}
