package Lem

import (
	"fmt"
	"strconv"
)

func Print(MesingPath *[][]string, a *info) {
	LenPath := make(map[int]int) // path and len of path

	for i := 0; i < len(*MesingPath); i++ {
		LenPath[i] = len((*MesingPath)[i])
	}

	PositionAnts := make([][]string, len(*MesingPath))
	var SmalPath int
	var Position int
	for i := 1; i <= a.nml; i++ {
		for j := 0; j < len(*MesingPath); j++ {
			if j == 0 {
				SmalPath = LenPath[0]
				Position = 0
			}
			if LenPath[j] < SmalPath {
				SmalPath = LenPath[j]
				Position = j
			}
		}
		PositionAnts[Position] = append(PositionAnts[Position], strconv.Itoa(i))
		LenPath[Position] = LenPath[Position] + 1
	}

	Ants_path := make(map[string][]string) 
	for i := 0; i < len(PositionAnts); i++ {  
		for _, va := range PositionAnts[i] { 
			Ants_path[va] = (*MesingPath)[i]
		}
	}
	CheckAntsSameWay := make(map[int]bool)
	slayceofRoms := []string{}

	var StockString string 
	for {
		for i := 1; i <= a.nml; i++ {
			Ants := strconv.Itoa(i)
			if len(Ants_path[Ants]) != 0 {
				if !Chekslayce(slayceofRoms, Ants_path[Ants][0]) || a.end == Ants_path[Ants][0] {
					if Ants_path[Ants][0] == a.end {
						if CheckAntsSameWay[Find(PositionAnts, Ants)] {
							continue
						} else {
							CheckAntsSameWay[Find(PositionAnts, Ants)] = true
						}
					}
					slayceofRoms = append(slayceofRoms, Ants_path[Ants][0])
					if i == a.nml {
						StockString += "L" + Ants + "-" + Ants_path[Ants][0]
					} else {
						StockString += "L" + Ants + "-" + Ants_path[Ants][0] + " "
					}
					Ants_path[Ants] = Ants_path[Ants][1:]
				}
			}

		}
		for x := range CheckAntsSameWay {
			CheckAntsSameWay[x] = false
		}

		if StockString == "" {
			break
		}
		fmt.Println(StockString)
		StockString = ""
		slayceofRoms = nil
	}
}
