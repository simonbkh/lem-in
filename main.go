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
		fmt.Println("Invalid Argumant ")
		return
	}

	lem.Parsing(arg[0])
}