package Lem

import (
	"fmt"
	"slices"
)

// func Print(MesingPath *[][]string, a *info) {
// 	LenPath := make(map[int]int) // path and len of path

// 	for i := 0; i < len(*MesingPath); i++ {
// 		LenPath[i] = len((*MesingPath)[i])
// 	}

// 	PositionAnts := make([][]string, len(*MesingPath))
// 	var SmalPath int
// 	var Position int
// 	for i := 1; i <= a.nml; i++ {
// 		for j := 0; j < len(*MesingPath); j++ {
// 			if j == 0 {
// 				SmalPath = LenPath[0]
// 				Position = 0
// 			}
// 			if LenPath[j] < SmalPath {
// 				SmalPath = LenPath[j]
// 				Position = j
// 			}
// 		}
// 		PositionAnts[Position] = append(PositionAnts[Position], strconv.Itoa(i))
// 		LenPath[Position] = LenPath[Position] + 1
// 	}

// 	Ants_path := make(map[string][]string)
// 	for i := 0; i < len(PositionAnts); i++ {
// 		for _, va := range PositionAnts[i] {
// 			Ants_path[va] = (*MesingPath)[i]
// 		}
// 	}
// 	CheckAntsSameWay := make(map[int]bool)
// 	slayceofRoms := []string{}

// 	var StockString string
// 	for {
// 		for i := 1; i <= a.nml; i++ {
// 			Ants := strconv.Itoa(i)
// 			if len(Ants_path[Ants]) != 0 {
// 				if !Chekslayce(slayceofRoms, Ants_path[Ants][0]) || a.end == Ants_path[Ants][0] {
// 					if Ants_path[Ants][0] == a.end {
// 						if CheckAntsSameWay[Find(PositionAnts, Ants)] {
// 							continue
// 						} else {
// 							CheckAntsSameWay[Find(PositionAnts, Ants)] = true
// 						}
// 					}
// 					slayceofRoms = append(slayceofRoms, Ants_path[Ants][0])
// 					if i == a.nml {
// 						StockString += "L" + Ants + "-" + Ants_path[Ants][0]
// 					} else {
// 						StockString += "L" + Ants + "-" + Ants_path[Ants][0] + " "
// 					}
// 					Ants_path[Ants] = Ants_path[Ants][1:]
// 				}
// 			}

// 		}
// 		for x := range CheckAntsSameWay {
// 			CheckAntsSameWay[x] = false
// 		}

// 		if StockString == "" {
// 			break
// 		}
// 		fmt.Println(StockString)
// 		StockString = ""
// 		slayceofRoms = nil
// 	}
// }

type nmilaat struct {
	name int
	path []string
	room int
}

func Print(mok *info, paths [][]string) {
	var prt []nmilaat
	var p nmilaat
	pathCost := []int{}

	// Ensure paths is properly defined and populated
	for _, v := range paths {
		pathCost = append(pathCost, len(v))
	}

	for i := 1; i <= mok.nml; i++ {
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
	fmt.Println(prt)
	sl := []int{}
	done := 0
	for {
		if done == -1 {
			break
		}
		done = -1

		arr := make([]bool, len(paths))
		mp := make(map[string]bool)
		for i := range prt {
			if len(prt[i].path) != 0 {
				done = 0
			}
			if slices.Contains(sl, i) {
				continue
			}

			if len(prt[i].path) != 0 {
				if !mp[prt[i].path[0]] {
					mp[prt[i].path[0]] = true
					fmt.Printf("L%d-%v ", prt[i].name, prt[i].path[0])
					prt[i].path = prt[i].path[1:]
				} 
			} else {
		
					
				if !arr[prt[i].room]{
					arr[prt[i].room] = true
					fmt.Printf("L%d-%v ", prt[i].name, mok.end)
					if !slices.Contains(sl, i) {
						sl = append(sl, i)
						// prt = append(prt[:i], prt[i+1:]... )
					}
				}
				
			}
		}
		fmt.Println()

	}
}
