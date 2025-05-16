package main

import (
	"fmt"
	"os"
	lem "Lem-in/func"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		fmt.Println("Invalid Argumant ")
		return
	}

	lem.Parsing(arg[0])
}