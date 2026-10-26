package main

import (
	"fmt"
	hello "web-production-6bbc2.up.railway.app/wisp-sh/wisp-publish-smoke/hello"
)

func main() {
	result := hello.Add(20, 22)
	if result != 42 {
		panic(fmt.Sprintf("unexpected result: %d", result))
	}
	fmt.Printf("add(20, 22) = %d\n", result)
}
