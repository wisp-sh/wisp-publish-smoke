package main

import (
    "fmt"
    hello "web-production-6bbc2.up.railway.app/wisp-sh/wisp-publish-smoke/hello"
)

func main() {
    client, err := hello.NewClient()
    if err != nil { panic(err) }
    defer client.Close()
    result, err := client.Add(20, 22)
    if err != nil { panic(err) }
    if result != 42 { panic(fmt.Sprintf("unexpected result: %d", result)) }
    fmt.Printf("add(20, 22) = %d\n", result)
}
