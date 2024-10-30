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
	var stock string
	stockSl := [][]string{}
	// rooms := make(map[interface{}]bool)
	content, err := os.ReadFile(fileName)
	if err != nil || len(content) == 0 {
		fmt.Println("file dialk fih machkil a 3chiri")
		return
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	// st, fin := false, false

	for scanner.Scan() {
		if i == 0 {
			nmilat, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("ERROR: invalid data format")
				return
			}
			if nmilat >= 1 {
				mok.nml = nmilat
			} else {
				fmt.Println("ERROR: invalid data format")
				return
			}

			fmt.Println(mok.nml)
		}
		// if st && len(mok.start) == 0 {
		// 	if len(mok.start) == 0 {
		// 		ysf := strings.Fields(scanner.Text())
		// 		mok.start = ysf[0]
		// 	}
		// }
		if stock == "##start"  {
			ysf := strings.Fields(scanner.Text())
			fmt.Println(ysf)
			if len(ysf) != 3 {
				fmt.Println("ERROR: invalid data format")
				return
			}
			stockSl = append(stockSl, ysf)
			mok.start = stockSl[0][0]
		}
		if stock == "##end" {
			ysf := strings.Fields(scanner.Text())
			stockSl = append(stockSl, ysf)
			mok.end = stockSl[0][0]

		}
		if strings.HasPrefix(scanner.Text(), "#") {
			if scanner.Text() == "##start" {
				stock = scanner.Text()
				// st = true
			} else if scanner.Text() == "##end" {
				stock = scanner.Text()
				// fin = true
			}
		}

		i++
	}
	fmt.Println(mok.start)
	fmt.Println(mok.end)
}
