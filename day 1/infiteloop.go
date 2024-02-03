package main

import (
    "fmt"
    "time"
)

func main() {
    // Initialize a variable for incrementing
    i := 0

    // Create a channel to control the loop
    stop := make(chan bool)

    // Start a goroutine to increment i and print
    go func() {
        for {
            // Print the value of i
            fmt.Printf("Shyama %d\n", i)

            // Increment i
            i++

            // Introduce a delay
            time.Sleep(time.Second)

            // Check if the loop should stop
            select {
            case <-stop:
                return
            default:
            }
        }
    }()

    // Wait for a key press to stop the loop
    fmt.Println("Press any key to stop...")
    var input string
    fmt.Scanln(&input)

    // Send a signal to stop the loop
    stop <- true
}
