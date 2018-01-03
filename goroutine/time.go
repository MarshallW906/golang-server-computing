package main

import (
    "fmt"
    "time"
)

func main() {
    tick := time.Tick(100 * time.Millisecond)
    boom := time.After(500 * time.Millisecond)
    for {
        select {
        case t := <-tick:
            fmt.Println("tick.", t)
        case t := <-boom:
            fmt.Println("BOOM!")
            return
//      default:
//          fmt.Println("    .")
//          time.Sleep(50 * time.Millisecond)
        }
    }
}

