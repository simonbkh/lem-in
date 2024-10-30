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

var Link = make(map[string][]string)

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
			if strings.HasPrefix(scanner.Text(), "#") {
			if scanner.Text() == "##start" {
				stock = scanner.Text()
				continue
				// st = true
			} else if scanner.Text() == "##end" {
				stock = scanner.Text()
				continue
				// fin = true
			}
		}
		if i == 0 {
			nmilat, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("ERROR: invalid data formatt")
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

	
		if stock == "##start" && !strings.HasPrefix(scanner.Text(), "#") {
			ysf := strings.Fields(scanner.Text())
			if len(ysf) != 3 {
				fmt.Println("ERROR: invalid data format")
				return
			}
			stockSl = append(stockSl, ysf)
			if len(mok.start) == 0 {
				mok.start = ysf[0]
			}
			mok.start = stockSl[0][0]

		}
		if stock == "##end" {
			if len(mok.start) != 0 {
				ysf := strings.Fields(scanner.Text())
				if len(mok.end) == 0 && len(ysf) == 3 {
					mok.end = ysf[0]
				}
				if len(mok.end) == 3 {
					stockSl = append(stockSl, ysf)
				}
				if len(ysf) == 1 {
					lin := strings.Split(scanner.Text(), "-")
					if len(lin) == 2 {
						Link[lin[0]] = append(Link[lin[0]], lin[1])
					}else{
						fmt.Println("ERROR: invalid data format")
						return
					}

				}

			} else {
				fmt.Println("ERROR: invalid data format")
				return
			}
		}
	
		i++
	}
	fmt.Println(Link)
	fmt.Println(mok.start)
	fmt.Println(mok.end)
}
