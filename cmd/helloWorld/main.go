package main

import (
	"fmt"

	"github.com/broadinstitute/golang-project-template/pkg/util"
)

func main() {
	fmt.Println(hello())
	util.PrintTriple(3)
}

func hello() string {
	return "Hello, World!"
}
