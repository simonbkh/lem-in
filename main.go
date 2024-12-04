package main

import (
	"fmt"
	"os"
	lem "Lemmok/func"
)

func main() {
	// var mok info
	arg := os.Args[1:]
	if len(arg) != 1 {
		fmt.Println("khoya dak args raj3hom o raje3 rask meahom")
		return
	}

	lem.Parsing(arg[0])
}