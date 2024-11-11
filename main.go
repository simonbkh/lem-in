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
var mok info

func main() {

	arg := os.Args[1:]
	if len(arg) != 1 {
		fmt.Println("khoya dak args raj3hom o raje3 rask meahom")
		return
	}
	Parsing(arg[0], &mok)
	CheckPath()
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
			// fmt.Println(a.nml)
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

	// fmt.Println(Link)
	// fmt.Println(a.start)
	// fmt.Println(a.end)
	// fmt.Println(uniRooms)
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

func CheckPath() {
	var path [][]string
	// checkDfs(&path, mok.start, []string{})
	Bfs(&path)
	// fmt.Println(path)
}

var roms = map[string]bool{}

func checkDfs(path *[][]string, rom string, sl []string) {
	start := rom
	roms[start] = true
	end := mok.end
	sl = append(sl, start)
	if start == end {
		*path = append(*path, sl)
	} else {
		// if slices.Contains(Link[start], end){

		// }
		for _, v := range Link[start] {
			if !roms[v] {
				checkDfs(path, v, sl)
			}
		}
	}
	roms[start] = false
}


////////////////////////////////
func Bfs(path *[][]string, start string, sl []string) {
	// count := 0
	// start := str
	end := mok.end
	queue := []string{start}
	visited := map[string][]string{start: {"none"}}
	for len(queue) > 0 {
		current := queue[0]
		fmt.Println("kkkkk")
		queue = queue[1:]
		// if current == end {
		// 	break
		// }
		fmt.Println(Link[current], current)
		for _, v := range Link[current] {

			_, found := visited[v]
			if !found || v == end {
				fmt.Println("---", v)
				visited[v] = append(visited[v], current)
				queue = append(queue, v)
			}

		}
	}
	// *path = append(*path, )
	fmt.Println(visited)
	// backtrackPath(visited, end, path, []string{})
	// fmt.Println(path)
}

func backtrackPath(visited map[string][]string, end string, paths *[][]string, path []string) {
	current := end
	path = append(path, current)
	if current == "1" || (len(*paths)!=0 && slices.Contains(*paths[0],current)){
		*paths = append(*paths, path)
	} else {
		for _, v := range visited[current] {
			backtrackPath(visited, v, paths, path)
		}
	}
	fmt.Println(paths)
	//

	// reverse(path)
	// return path
}

func reverse(array []string) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}


///////////////////////////////////
// var (
// 	qeu     []string
// 	romlink = map[string][]string{}
// )
// ////////////////////////////
// func Bfs(path *[][]string) {
// 	// count := 0
// 	stare := mok.start
// 	end := mok.end
// 	roms[stare] = true
// 	count := Apendqeu(stare)
// 	fmt.Println(qeu)
// 	for _, v := range qeu {
// 		// if count != 0 {
// 			romlink[v] = append(romlink[v], qeu[0])
// 			count--
// 			fmt.Println(qeu[0], stare, "mmmm")
// 		// 	// fmt.Println("kkk",stare,v,romlink)
// 		// }
// 		if slices.Contains(Link[v], end) {
// 			// index := slices.Index(Link[v], end)
// 			romlink[v] = append(romlink[v], end)
// 			qeu = qeu[1:]

// 			continue
// 			// if index != len(qeu)-1 {
// 			// 	qeu = append(qeu[:index], qeu[index+1:]...)
// 			// }

// 		} else {

// 			Apendqeu(v)
// 			fmt.Println("kkk", qeu, v)
// 			qeu = qeu[1:]
// 			fmt.Println("kkk111", qeu, v, roms)
// 		}
// 		// romlink[v] = append(romlink[v], end)
// 		//
// 	}
// 	fmt.Println(romlink)
// }

// func Rmovqeu(index int) {
// 	qeu = qeu[1:]
// 	qeu = append(qeu[:index], qeu[index+1:]...)
// }

// func Apendqeu(v string) int {
// 	conut := 0
// 	for _, c := range Link[v] {
// 		// fmt.Println("hhhhhhhhhh", c,Link[v],v)
// 		conut++
// 		if !roms[c] {
// 			qeu = append(qeu, c)
// 			roms[c] = true
// 		}

// 	}

// 	return conut
// }
