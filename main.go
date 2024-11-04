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
var Point = make(map[string]bool)
var Cord = make(map[string]bool)

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
				i++
				continue
			} else if scanner.Text() == "##end" {
				stock = scanner.Text()
				i++
				continue
			} else {
				i++
				continue
			}
		}
		if i == 0 {
			nmilat, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
				return
			}
			if nmilat >= 1 {
				mok.nml = nmilat
			} else {
				fmt.Println("ERROR: invalid data format")
				return
			}
			i++
			continue
		}
		if stock == "" && i != 0 {
			ysf := strings.Fields(scanner.Text())
			if len(ysf) != 3 {
				fmt.Println("ERROR: invalid data format")
				return
			} else if len(ysf) == 3 && !Point[ysf[0]] && !Cord[ysf[1]+" "+ysf[2]] {
				Point[ysf[0]] = true
			} else {
				fmt.Println("ERROR: invalid data format")
				return
			}

		}

		if stock == "##start" {

			ysf := strings.Fields(scanner.Text())
			if len(ysf) != 3 {
				fmt.Println("ERROR: invalid data format")
				return
			}

			if len(ysf) == 3 {
				if len(mok.start) == 0 {
					mok.start = ysf[0]
				}
				check(ysf)
			}

		}
		if stock == "##end" && mok.start != "" {
			ysf := strings.Fields(scanner.Text())
			if mok.end == "" && len(ysf) == 3 {
				mok.end = ysf[0]
			}
			if len(ysf) == 3 {
				check(ysf)
			}
			if len(ysf) == 1 {
				lin := strings.Split(scanner.Text(), "-")
				if len(lin) == 2 {
					Link[lin[0]] = append(Link[lin[0]], lin[1])
					Link[lin[1]] = append(Link[lin[1]], lin[0])
				} else {
					fmt.Println("ERROR: invalid data format")
					return
				}

			}

		} else if stock == "##end" && mok.start != "" {
			fmt.Println("ERROR: invalid data format")
			return
		}
		i++
	}
	fmt.Println(mok.start)
	fmt.Println(mok.end)
	fmt.Println(Link)
	BFS()
}

func check(ysf []string) {
	for _, Numbre := range ysf[1:] {
		nb, err := strconv.Atoi(Numbre)
		if err != nil || nb < 0 {
			fmt.Println("ERROR: invalid data format")
			os.Exit(1)
		}
	}
	if !Point[ysf[0]] {
		Point[ysf[0]] = true
		Cord[ysf[1]+" "+ysf[2]] = true
	} else {
		fmt.Println("ERROR: invalid data format")
		os.Exit(1)
	}
}


func BFS() {
	slayce := [][]string{}
	Taak := make(map[string]bool)
	albdya := mok.start
	sl := []string{albdya}
	for len(sl)!=0{
		albdya = sl[0]
		if !Taak[albdya]{
			Taak[albdya]=true
			fmt.Println("hio")
		for _,val := range Link[albdya]{
			sl = append(sl, val)
		}
		
		}
		sl = sl[1:]
		fmt.Println(sl)
	}

}
