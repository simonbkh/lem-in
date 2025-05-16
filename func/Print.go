package Lem

import (
	"fmt"
	"slices"
)

type nmilaat struct {
	name int
	path []string
	room int
}

func Print(inf *info, paths [][]string) {
	var prt []nmilaat
	var p nmilaat
	pathCost := []int{}

	// Ensure paths is properly defined and populated
	for _, v := range paths {
		pathCost = append(pathCost, len(v))
	}

	for i := 1; i <= inf.nml; i++ {
		p.name = i
		if len(pathCost) == 0 {
			return // Prevent panic if pathCost is empty
		}
		ii := slices.Index(pathCost, slices.Min(pathCost))
		pathCost[ii]++
		p.path = paths[ii][:len(paths[ii])-1]
		p.room = ii
		prt = append(prt, p)
		p = nmilaat{}
	}

	sl := []int{}
	// done := 0
	s := ""
	for {
		// if done == -1 {
		// 	break
		// }
		// done = -1

		arr := make([]bool, len(paths))
		mp := make(map[string]bool)
		for i := range prt {

			if slices.Contains(sl, i) {
				continue
			}

			if len(prt[i].path) != 0 {
				// done = 0
				if !mp[prt[i].path[0]] {
					mp[prt[i].path[0]] = true
					s += fmt.Sprintf("L%d-%v ", prt[i].name, prt[i].path[0])
					prt[i].path = prt[i].path[1:]
				}
			} else {
				if !arr[prt[i].room] {
					arr[prt[i].room] = true
					s += fmt.Sprintf("L%d-%v ", prt[i].name, inf.end)
					sl = append(sl, i)

				}
			}
		}
		if s == "" {
			break
		}
		fmt.Println(s[:len(s)-1])
		s = ""

	}
}
