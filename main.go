package main

import (
	"bufio"
	"fmt"
	"os"
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
	// stockSl := [][]string{}
	//cords := make(map[string]bool)
	uniRooms := make(map[string]bool)
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
				a.nml = nmilat
			} else {
				fmt.Println("ERROR: invalid data format")
				return
			}

			fmt.Println(a.nml)
			i++
			continue
		}
		if strings.HasPrefix(scanner.Text(), "#") {
			if scanner.Text() == "##start" {
				//fmt.Println(stock)
				if !st {
					st = true
					continue
				} else {
					fmt.Println("ERROR9: invalid data format")
					return
				}

				// st = true
			} else if scanner.Text() == "##end" {
				if !fin {
					fin = true
					continue
				} else {
					fmt.Println("ERROR*: invalid data format")
					return
				}
				// fin = true
			}
			continue
		}
		ysf := strings.Fields(scanner.Text())
		if scanner.Text() != "" {

			// if len(ysf) != 3 {
			// fmt.Println(ysf)
			// 	return
			// }else{
			if len(ysf) == 3 {
				if !uniRooms[ysf[0]] {
					uniRooms[ysf[0]] = true
				} else {
					fmt.Println("room meawda a 3chiri")
					return
				}
			} else if len(ysf) == 1 && st && fin{
				lin := strings.Split(scanner.Text(), "-")
				if len(lin) == 2 {
					if uniRooms[lin[0]] {
						Link[lin[0]] = append(Link[lin[0]], lin[1])
					} else {
						fmt.Println(scanner.Text())
						fmt.Println("ERROR: room mam3rofach")
						return
					}

				} else {
					fmt.Println("ERROR$$: invalid data format")
					return
				}

			}else{
				fmt.Println("ERROR: invalid data format")
				return
			}
		} else {
			fmt.Println("ERROR^: invalid data format")
			return
		}
		if st && len(a.start) == 0 {
			// stockSl = append(stockSl, ysf)
			if len(ysf) != 3 {
				fmt.Println("ERROR': invalid data format")
				return
			}
			if len(a.start) == 0 {
				a.start = ysf[0]
			}
			// mok.start = stockSl[0][0]
		}
		if fin && len(a.end) == 0 {
			if len(ysf) == 3 {
				a.end = ysf[0]
			}

		}
	}

	if len(a.end) == 0 || len(a.start) == 0 {
		fmt.Println("ERROR: invalid data format")
		return
	}
	//endc , startc := false,false
	//_, startc := Link[a.start]
	
	if !(check(a.end) && check(a.start)) {
		fmt.Println("azeby start awla end ra mamlinkinch")
		return
	}
	
	

	fmt.Println(Link)
	fmt.Println(a.start)
	fmt.Println(a.end)
	fmt.Println(uniRooms)

}
func check(s string)bool{
	_, ok :=  Link[s]
	if  !ok{
		for _, v := range Link {
			for _, mok := range v {
				if mok == s {
					return true
				}
			}
	   }
	}else{
		return true
	}
	return false
}
