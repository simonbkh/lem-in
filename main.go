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

func Print(Path *[][]string, a *info) {
	mapp := make(map[int]int)
	for i := 0; i < len(*Path); i++ {
		mapp[i] = len((*Path)[i])
	}

	sl := make([][]string, len(*Path))
	var sml int
	var mapsmal int
	for i := 1; i <= a.nml; i++ {
		for j := 0; j < len(*Path); j++ {
			if j == 0 {
				sml = mapp[0]
				mapsmal = 0
			}
			if mapp[j] < sml {
				sml = mapp[j]
				mapsmal = j
			}
		}
		sl[mapsmal] = append(sl[mapsmal], strconv.Itoa(i))
		mapp[mapsmal] = mapp[mapsmal] + 1
	}
	fmt.Println(len(mapp))
	fmt.Println(mapp)
	fmt.Println(sl)
	 Mapcheck := make(map[string][]string)
	for i := 0; i < len(sl); i++ {
		for _ , va := range sl[i] {
			 Mapcheck[va] = (*Path)[i]
		}
		
	}
	fmt.Println(Mapcheck)
	sla := []string{}
	var s string
	for {
		for key , valeu := range Mapcheck {
			if len(valeu) !=0 {
				if !Chekslayce(sla, valeu[0]) {
					sla = append(sla,valeu[0])
					s += "L"+key+"-"+valeu[0]+" "
					Mapcheck[key] = valeu[1:]
				}
			}
		}
		if s == "" {
			break
		}
		fmt.Println(s)
		s = ""
		sla = nil
	}
	

	
}

func Chekslayce(Path []string, s string, )bool {
	for _, v := range Path {
		if v == s {
			return true
		}
		// if v == *a {
		// 	return false
		// }
	}
	return false

}

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
	Print(&m, a)
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
			pp = append(pp, paths[index][1:])
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
