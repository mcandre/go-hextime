// Package main provides a hextime executable.
package main

import (
	"fmt"

	"github.com/mcandre/go-hextime"
)

// main is the entrypoint for launching this application.
func main() {
	fmt.Println(hextime.Hextime())
}
