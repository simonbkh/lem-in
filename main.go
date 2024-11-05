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

var (
	Link = make(map[string][]string)
)

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
				if !st && len(a.end) == 0 && !fin{
					st = true
					continue
				} else {
					fmt.Println("ERROR9: invalid data format", scanner.Text())
					return
				}

			} else if scanner.Text() == "##end" {
				if !fin && len(a.start) != 0 {
					fin = true
					continue
				} else {
					fmt.Println("ERROR*: invalid data format", scanner.Text())
					return
				}
			}
			continue
		}
		ysf := strings.Fields(scanner.Text())
		if scanner.Text() != "" {
			if len(ysf) == 3 {
				if !uniRooms[ysf[0]] {
					uniRooms[ysf[0]] = true
				} else {
					fmt.Println("room meawda a 3chiri")
					return
				}
			} else if len(ysf) == 1 && st && fin {
				lin := strings.Split(scanner.Text(), "-")
				if len(lin) == 2 {
					if uniRooms[lin[0]] &&   uniRooms[lin[1]]{ // nfekro flm3awda!
						Link[lin[0]] = append(Link[lin[0]], lin[1])
						Link[lin[1]] = append(Link[lin[1]], lin[0])

					} else {
						fmt.Println("ERROR: room mam3rofach", scanner.Text())
						return
					}

				} else {
					fmt.Println("ERROR$$: invalid data format")
					return
				}

			} else {
				fmt.Println("ERROR: invalid data format")
				return
			}
		} else {
			fmt.Println("ERROR^: invalid data format")
			return
		}
		if st && len(a.start) == 0 {
			if len(ysf) != 3 {
				fmt.Println("ERROR': invalid data format")
				return
			}
			if len(a.start) == 0 {
				a.start = ysf[0]
			}
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

	if !(check(a.end) && check(a.start)) {
		fmt.Println("khoya!! start awla end ra mamlinkinch")
		return
	}

	fmt.Println(Link)
	p := findAllPaths(a.start, a.end)
	fmt.Println(p)

	m := MesingPath(p)
	fmt.Println(m)

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

func findAllPaths(start, end string) [][]string {
	var paths [][]string
	visited := make(map[string]bool)
	dfs(start, end, visited, []string{}, &paths)
	return paths
}

func dfs(start, end string, visited map[string]bool, currentPath []string, paths *[][]string) {
	visited[start] = true
	currentPath = append(currentPath, start)
	// fmt.Println(start,currentPath)
	// fmt.Println(visited)
	//fmt.Println(currentPath)

	if start == end {
		*paths = append(*paths, append([]string{}, currentPath[1:]...))
	} else {
		for _, neighbor := range Link[start] {

			if !visited[neighbor] {
				dfs(neighbor, end, visited, currentPath, paths)
			}
			if (neighbor==end){
				break
			}
		}
	}

	visited[start] = false
	//currentPath = currentPath[:len(currentPath)-1]
}


func MesingPath(paths [][]string) []int {
	var pp [][]string
	var p []int
	var Nber int
	for _, v := range paths {
		for _, i := range v {
			for _, j := range paths {
				for _, k := range j {
					if i == k {
						Nber++
					}
				}
				
			}
			
		}
		p = append(p, Nber)
		Nber = 0
	}

	min, index := Asrar(p)
	pp =append(pp, paths[index])

	for _, v := range paths {
		for _, i := range v {
			
		}
	}
	return p
}

func Asrar(p []int) (int, int) {
	var min int
	var index int
	min = p[0]
	for r, v := range p {
		if v < min {
			min = v
			index = r
		}
	}
	return min , index
}