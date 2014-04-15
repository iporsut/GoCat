package main

import (
    "log"
    "io"
    "os"
)

func main() {
    stdin := os.Stdin
    if len(os.Args) > 1 {
        file, err := os.Open(os.Args[1])
        if err != nil {
            log.Fatal(err)
        } else {
            stdin = file
        }
    }
    io.Copy(os.Stdout, stdin)
}
