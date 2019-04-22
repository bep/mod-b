package main

import (
	"fmt"
)

var (
	name    = "b"
	version = "v1.1"
)

func main() {
	fmt.Println(Version())
}

func Version() string {
	return name + ": " + version
}
