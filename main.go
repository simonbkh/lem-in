package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var mok info

type info struct {
	nml   int
	start string
	end   string
}

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		fmt.Println("khoya dak args raj3hom o raje3 rask meahom")
		return
	}
	Parsing(arg[0])
}

func Parsing(fileName string) {
	//rooms := make(map[interface{}]bool)
	content, err := os.ReadFile(fileName)
	if err != nil || len(content) == 0 {
		fmt.Println("file dialk fih machkil a 3chiri")
		return
	}
	// contSplited := strings.Split(string(content), string('\n'))

	// for i, line := range contSplited {
	// 	if i == 0 {
	// 		nmilat, err := strconv.Atoi(line)
	// 		if err != nil {
	// 			fmt.Println("ERROR: invalid data format")
	// 			return
	// 		}
	// 		if nmilat >= 1 {
	// 			mok.nml = nmilat
	// 		} else {
	// 			fmt.Println("ERROR: invalid data format")
	// 			return
	// 		}
	// 	}
	// }
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	st, fin := false, false

	for scanner.Scan() {
		if i == 0 {
			nmilat, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("ERROR: invalid data format")
				return
			}
			if nmilat >= 1 {
				mok.nml = nmilat
			} else if nmilat == 0{
				return
			}else{
				fmt.Println("ERROR: invalid data format")
				return
			}
			fmt.Println(mok.nml)
		}
		if st{
			if len(mok.start) == 0 {
				ysf := strings.Split(scanner.Text(), " ")
				
				mok.start = ysf[0]
			}
			

		}
		if fin  && len(mok.end) == 0 {
			
				ysf := strings.Split(scanner.Text(), " ")
				mok.end = ysf[0]
	
		}
		if strings.HasPrefix(scanner.Text(), "#") {
			if scanner.Text() == "##start" {
				st = true
			} else if scanner.Text() == "##end" {
				fin = true
			}
		}

		i++
	}
	fmt.Println(mok.start)
	fmt.Println(mok.end)
}
