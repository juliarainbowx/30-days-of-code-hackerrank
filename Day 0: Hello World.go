package main

import (
    "fmt"
    "os"
    "io"
)

func main() {
 fmt.Println("Hello, World.")
 io.Copy(os.Stdout, os.Stdin)
}
