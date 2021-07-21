package main

import (
	"fmt"

	"github.com/broadinstitute/golang-project-template/internal/util"
)

func main() {
	fmt.Println(hello())
	util.PrintTriple(3)
}

// demo function
func hello() string {
	return "Hello, World!"
}

func doSomething() {
	// this is a do nothing function added purely for demonstration
}
