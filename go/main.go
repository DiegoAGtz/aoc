package main

import (
	"fmt"
	"os"
)

func main() {
    data, err := os.ReadFile("input.txt")
    check(err)
    floor := 0

    for i := 0; i < len(data); i++ {
        if data[i] == '(' {
            floor++
        } else {
            floor -= 1
        }
        if floor == -1 {
            fmt.Println(i)
        }
    }
    fmt.Println(len(data), floor)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
