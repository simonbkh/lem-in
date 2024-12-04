package Lem

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

func Parsing(fileName string) {
	var st, fin bool
	var a info
	uniRooms := make(map[string]bool)
	Rooms := make(map[string]string)
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
		room := strings.Fields(scanner.Text())
		if scanner.Text() != "" {
			if len(room) == 3 {
				if !uniRooms[room[0]] && !strings.HasPrefix(room[0], "L") {
					_, err := strconv.Atoi(room[1])
					_, er := strconv.Atoi(room[2])
					if err == nil && er == nil {
						uniRooms[room[0]] = true
						if _, ok := Rooms[strings.Join(room[1:], " ")]; !ok {
							Rooms[strings.Join(room[1:], " ")] = room[0]
						} else {
							fmt.Println("ERROR: invalid data format, invalid coordinates")
							return
						}
					} else {
						fmt.Println("ERROR: invalid data format, invalid coordinates")
						return
					}

				} else {
					fmt.Println("ERROR: invalid data format, invalid Rooms")
					return
				}
			} else if len(room) == 1 && st && fin {
				lin := strings.Split(scanner.Text(), "-")
				if len(lin) == 2 {
					if uniRooms[lin[0]] && uniRooms[lin[1]] && lin[0] != lin[1] {
						ind := slices.Index(Link[lin[0]], lin[1])
						if ind == -1 {
							Link[lin[0]] = append(Link[lin[0]], lin[1])
							Link[lin[1]] = append(Link[lin[1]], lin[0])
						} else {
							fmt.Println("ERROR: invalid data format, repeated Link")
							return
						}

					} else {
						fmt.Println("ERROR: invalid data format, invalid Link")
						return
					}
				} else {
					fmt.Println("ERROR: invalid data format, invalid Link")
					return
				}

			} else {
				fmt.Println("ERROR: invalid data format")
				return
			}
		} else {
			fmt.Println("ERROR: invalid data format")
			return
		}
		if st && len(a.start) == 0 {
			if len(room) != 3 {
				fmt.Println("ERROR': invalid data format")
				return
			}
			if len(a.start) == 0 {
				a.start = room[0]
			}
		}
		if fin && len(a.end) == 0 {
			if len(room) == 3 {
				a.end = room[0]
			}
		}
	}

	if len(a.end) == 0 || len(a.start) == 0 {
		fmt.Println("ERROR: invalid data format")
		return
	}

	if len(a.end) == 0 && len(a.start) == 0 {
		fmt.Println("ERROR, invalid data format start or end rooms aren't connected")
		return
	}
	p := findAllPaths(&a)

	MesingPath := MesingPath(p)
	fil, errr := os.ReadFile("test/t00.txt")
	if errr != nil {
		fmt.Fprintln(os.Stderr, "ERROR reading file")
	}
	fmt.Println(string(fil) + "\n")
	if len(MesingPath) != 0 {
		Print(&a, MesingPath)
	} else {
		fmt.Fprintln(os.Stderr, "no paths found")
		return
	}
}
