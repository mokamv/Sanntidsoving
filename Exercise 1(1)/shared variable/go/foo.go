// Use `go run foo.go` to run your program

package main

import (
	. "fmt"
	"runtime"
	"time"
)

var i = 0

func incrementing(c_inc chan bool) {
    for j:= 0; j < 1000000; j++ {
        c_inc <- true
    }
    close(c_inc)
}

func decrementing(c_dec chan bool) {
    for j:= 0; j < 1000000; j++ {
        c_dec <- true
    }
    close(c_dec)
}

func server(c_inc chan bool, c_dec chan bool, c_done chan bool) {
    for {
        select {
        case _, ok := <-c_inc:
            if ok {
                i += 1
            } else {
                c_inc = nil
            }
        case _, ok := <-c_dec:
            if ok {
                i -= 1
            } else {
                c_dec = nil
            }
        }
        if c_inc == nil && c_dec == nil {
            break
        }
    }
    c_done <- true
}

func main() {
    // What does GOMAXPROCS do? What happens if you set it to 1?
    runtime.GOMAXPROCS(2)

    c_inc := make(chan bool)
    c_dec := make(chan bool)
    c_done := make(chan bool)
	
    // TODO: Spawn both functions as goroutines
    go incrementing(c_inc)
    go decrementing(c_dec)
	go server(c_inc, c_dec, c_done)

    // We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
    // We will do it properly with channels soon. For now: Sleep.
    <-c_done // Venter med å kjøre main til en verdi er sendt til c_done
    time.Sleep(500*time.Millisecond)
    Println("The magic number is:", i)
}
