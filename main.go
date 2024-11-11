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

type nmilaat struct {
	name int
	path []string
	room string
}

var prt []nmilaat

func printer(mok *info, mat [][]string) {
	mp := make(map[string]bool)
	var p nmilaat
	pathCost := []int{}
	for _, v := range mat {
		pathCost = append(pathCost, len(v[1:]))
	}
	for i := 1; i <= mok.nml; i++ {
		p.name = i
		ii := slices.Index(pathCost, slices.Min(pathCost))
		pathCost[ii]++
		p.path = mat[ii][1:]
		prt = append(prt, p)
		p = nmilaat{}

	}
	for {
		for _, nmla := range prt {
			if len(nmla.path) != 0 {
				if nmla.room == mok.end {
					fmt.Printf("L%d-%v ", nmla.name, mok.end)
					nmla.path = []string{}
					continue
				}
				if mp[nmla.room] {
					mp[nmla.room] = false
				} else {
					fmt.Printf("L%d-%v ", nmla.name, nmla.path[0])
					nmla.room = nmla.path[0]
					mp[nmla.path[0]] = true
					nmla.path = nmla.path[0:]
				}

			}
		}
	}
	//fmt.Println(prt)
	// fmt.Println(pathCost)
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
				if !st && len(a.end) == 0 && !fin {
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
					if uniRooms[lin[0]] && uniRooms[lin[1]] { // nfekro flm3awda!
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

	if len(a.end) == 0 && len(a.start) == 0 {
		fmt.Println("khoya!! start awla end ra mamlinkinch")
		return
	}

	fmt.Println(Link)
	p := findAllPaths(a)
	fmt.Println(p)

	m := MesingPath(p)
	fmt.Println(m)
	printer(a, m)
}

func findAllPaths(m *info) [][]string {
	var paths [][]string
	visited := make(map[string]bool)
	dfs(m.start, m.end, visited, []string{}, &paths, m)
	return paths
}

func dfs(start, end string, visited map[string]bool, currentPath []string, paths *[][]string, m *info) {
	visited[start] = true
	currentPath = append(currentPath, start)

	if start == end {
		*paths = append(*paths, append([]string{}, currentPath...))
	} else {
		good := false
		if start != m.start {
			for _, v := range Link[start] {
				if v == end {
					good = true
					currentPath = append(currentPath, v)
					*paths = append(*paths, append([]string{}, currentPath...))
					break
				}
			}
		}
		if !good {
			for _, neighbor := range Link[start] {
				if !visited[neighbor] {
					dfs(neighbor, end, visited, currentPath, paths, m)
				}
				// if neighbor == end {
				// 	break
				// }
			}
		}

	}

	visited[start] = false
	// currentPath = currentPath[:len(currentPath)-1]
}

func MesingPath(paths [][]string) [][]string {
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

	mpp := make(map[string]bool)
	for i := 0; i < len(p); i++ {
		index := Small(&p)

		if check(paths[index][1:len(paths[index])-1], &mpp) {
			pp = append(pp, paths[index])
		}
	}

	// pp = append(pp, paths[index])
	return pp
}

func Small(p *[]int) int {
	var min int
	var index int
	min = (*p)[0]
	for r, v := range *p {
		if min == -1 && v != -1 {
			min = v
			index = r
		}
		if v == -1 {
			continue
		}
		if v < min {
			min = v
			index = r
		}
	}
	(*p)[index] = -1
	// fmt.Println(*p)
	return index
}

func check(path []string, mp *map[string]bool) bool {
	temp := []string{}
	for _, room := range path {
		if !(*mp)[room] {
			(*mp)[room] = true
			temp = append(temp, room)
		} else {
			for _, v := range temp {
				delete((*mp), v)
			}
			fmt.Println("false")
			return false
		}
	}
	fmt.Println("true")
	return true
}
