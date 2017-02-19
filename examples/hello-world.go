package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Hello There!")

    fmt.Println(time.Now().Format("2006-01-02_15_04_05"))
    time.Sleep(time.Second)
    fmt.Println(time.Now().Format("2006-01-02_15_04_05"))
}