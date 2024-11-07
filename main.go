package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type info struct {
	nml   int
	start string
	end   string
}

var Link = make(map[string][]string)

func main() {
	var mok info
	arg := os.Args[1:]
	if len(arg) != 1 {
		fmt.Println("khoya dak args raj3hom o raje3 rask meahom")
		return
	}
	Parsing(arg[0], &mok)
}

func Parsing(fileName string, a *info) {
	var st, fin bool
	var conut int
	// stockSl := [][]string{}
	// cords := make(map[string]bool)
	uniRooms := make(map[string]bool)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		if i == 0 {
			nmilat, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("ERROR: invalid data format")
				return
			}
			if nmilat < 1 {
				fmt.Println("ERROR: invalid data format")
				return
			}
			a.nml = nmilat
			fmt.Println(a.nml)
			i++
			continue
		}
		if strings.HasPrefix(scanner.Text(), "#") {
			if scanner.Text() == "##start" {
				if st {
					fmt.Println("ERROR9: invalid data format")
					return
				}
				st = true
				conut++
				continue
			} else if scanner.Text() == "##end" {
				if fin {
					fmt.Println("ERROR*: invalid data format")
					return
				}
				fin = true
				conut++
				continue
			}
			continue
		}
		ysf := strings.Fields(scanner.Text())
		if scanner.Text() != "" {
			if len(ysf) == 3 {
				if !CheckCordone(ysf) {
					fmt.Println("cordone khata ")
					return
				}
				if !uniRooms[ysf[0]] && !(strings.HasPrefix(ysf[0], "L")) {
					uniRooms[ysf[0]] = true
				} else {
					fmt.Println("room meawda a 3chiri")
					return
				}
			} else if len(ysf) == 1 && st && fin {
				lin := strings.Split(scanner.Text(), "-")
				if len(lin) != 2 {
					fmt.Println("ERROR$$: invalid data format")
					return
				}
				if !(uniRooms[lin[0]] && uniRooms[lin[1]]) {
					fmt.Println("ERROR: room mam3rofach")
					return
				}
				if slices.Contains(Link[lin[1]], lin[0]) {
					fmt.Println("had ruot rah m3awd", lin[0])
					return
				}
				Link[lin[0]] = append(Link[lin[0]], lin[1])

			} else {
				fmt.Println("ERROR: invalid data format")
				return
			}
		} else {
			fmt.Println("ERROR^: invalid data format")
			return
		}
		if st && len(a.start) == 0 {
			if len(ysf) != 3 || conut != 1 {
				fmt.Println("ERROR'333: invalid data format")
				return
			}
			if len(a.start) == 0 {
				a.start = ysf[0]
				conut--
			}
		}
		if fin && len(a.end) == 0 {
			if len(ysf) == 3 || conut != 1 {
				a.end = ysf[0]
				conut--
			}
		}
	}

	if len(a.end) == 0 || len(a.start) == 0 {
		fmt.Println("ERROR: invalid data format")
		return
	}
	if !(check(a.end) && check(a.start)) {
		fmt.Println("azeby start awla end ra mamlinkinch")
		return
	}

	fmt.Println(Link)
	fmt.Println(a.start)
	fmt.Println(a.end)
	fmt.Println(uniRooms)
}

func check(s string) bool {
	_, ok := Link[s]
	if !ok {
		for _, v := range Link {
			for _, mok := range v {
				if mok == s {
					return true
				}
			}
		}
	} else {
		return true
	}
	return false
}

func CheckCordone(sl []string) bool {
	for i := 1; i < len(sl); i++ {
		_, err := strconv.Atoi(sl[i])
		if err != nil {
			return false
		}
	}
	return true
}

func Checkdfs() {
}
