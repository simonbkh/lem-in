package Lem

import (
	"fmt"
	"strconv"
)

func Print(Path *[][]string, a *info) {
	LenPath := make(map[int]int)

	for i := 0; i < len(*Path); i++ { // had lop drtha bax n7at lin dyal kola path
		LenPath[i] = len((*Path)[i])
	}

	PositionAnts := make([][]string, len(*Path)) //
	var SmallerLenthPath int                     ///had variable drto bax n7dd bih bih asra4 len(dyal path)
	var AntsPosition int                         /// had variable drto bax n7dd bih alblasa axmn path atmchi fih nmla dyali
	for i := 1; i <= a.nml; i++ {                //// f had lop kan7t kola Ants f axmn tri9 atkon
		for j := 0; j < len(*Path); j++ {
			if j == 0 {
				SmallerLenthPath = LenPath[0]
				AntsPosition = 0
			}
			if LenPath[j] < SmallerLenthPath {
				SmallerLenthPath = LenPath[j]
				AntsPosition = j
			}
		}
		PositionAnts[AntsPosition] = append(PositionAnts[AntsPosition], strconv.Itoa(i))
		LenPath[AntsPosition] = LenPath[AntsPosition] + 1
	}

	CheckAntsSameWay := make(map[int]bool)
	for x := range *Path {
		CheckAntsSameWay[x] = false
	}

	KantsAndVpath := make(map[string][]string) ////had almap kan7at fih path dyal kola nmla key how ants o path how valeu
	for i := 0; i < len(PositionAnts); i++ {   ////kanlopi 3la path dyal kola wa7da o kan3tald 3la index bax n7at path alkola wa7da
		for _, va := range PositionAnts[i] { ////mital [[1 4 7 10] [2 5 8] [3 6 9]]
			KantsAndVpath[va] = (*Path)[i]
		}
	}

	slayceofRoms := []string{}

	var StockString string // had var bax n5zn tour L1-E L2-A L3-o L4-t L5-h L6-0
	for {
		for i := 1; i <= a.nml; i++ {
			Ants := strconv.Itoa(i)
			if len(KantsAndVpath[Ants]) != 0 {
				if !Chekslayce(slayceofRoms, KantsAndVpath[Ants][0]) || a.end == KantsAndVpath[Ants][0] {
					if KantsAndVpath[Ants][0] == a.end {
						if CheckAntsSameWay[Find(PositionAnts, Ants)] {
							continue
						} else {
							CheckAntsSameWay[Find(PositionAnts, Ants)] = true
						}
					}
					slayceofRoms = append(slayceofRoms, KantsAndVpath[Ants][0])
					if i == a.nml {
						StockString += "L" + Ants + "-" + KantsAndVpath[Ants][0]
					} else {
						StockString += "L" + Ants + "-" + KantsAndVpath[Ants][0] + " "
					}
					KantsAndVpath[Ants] = KantsAndVpath[Ants][1:]
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
