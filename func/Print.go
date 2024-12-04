package Lem

import (
	"fmt"
	"slices"
	"strings"
)


type nmilaat struct {
	name int
	path []string
	ind int
}

func Print(mok *info, paths [][]string) {
	var antData []nmilaat
	var temp nmilaat
	pathCost := []int{}

	// Ensure paths is properly defined and populated
	for _, v := range paths {
		pathCost = append(pathCost, len(v))
	}

	for i := 1; i <= mok.nml; i++ {
		temp.name = i
		ii := slices.Index(pathCost, slices.Min(pathCost))
		pathCost[ii]++
		temp.path = paths[ii][:len(paths[ii])-1]
		temp.ind = ii
		antData = append(antData, temp)
		temp = nmilaat{}
	}
	skippable := []int{}

	for {
		s := ""
	

		arr := make([]bool, len(paths))
		mp := make(map[string]bool)
		for i := range antData {
		
			if slices.Contains(skippable, i) {
				continue
			}

			if len(antData[i].path) != 0 {
				if !mp[antData[i].path[0]] {
					mp[antData[i].path[0]] = true
					s +=fmt.Sprintf("L%d-%v ", antData[i].name, antData[i].path[0])
					antData[i].path = antData[i].path[1:]
				} 
			} else {	
				if !arr[antData[i].ind]{
					arr[antData[i].ind] = true
					s +=fmt.Sprintf("L%d-%v ", antData[i].name, mok.end)
					if !slices.Contains(skippable, i) {
						skippable = append(skippable, i)
					}
				}
				
			}
		}
		if s == "" {
			break
		}
		fmt.Println(strings.TrimSpace(s))
		s = ""
	}
}
