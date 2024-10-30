package main

import (
    "errors"
    "fmt"
)

// Custom function that creates and wraps an error
func readFile(filename string) error {
    // Simulate an error (e.g., file not found)
    baseErr := errors.New("file not found")

    // Wrap the base error with more context using %w
    return fmt.Errorf("failed to read file %s: %w", filename, baseErr)
}

func main() {
    err := readFile("example.txt")
    if err != nil {
        fmt.Println("Error:", err)

        // Unwrapping the error to check the base error
        if errors.Is(err, errors.New("file not found")) {
            fmt.Println("The file was not found.")
        }
    }
}
