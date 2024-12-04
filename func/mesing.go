package Lem

func MesingPath(paths [][]string) [][]string {
	var MesingPath [][]string
	var ScorofPathe []int
	var Nber int
	// fmt.Println(paths)
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
		ScorofPathe = append(ScorofPathe, Nber)
		Nber = 0
	}

	CheckVisited := make(map[string]bool)
	for i := 0; i < len(ScorofPathe); i++ {
		index := Small(&ScorofPathe)

		if check(paths[index][1:len(paths[index])-1], &CheckVisited) {
			MesingPath = append(MesingPath, paths[index][1:])
		}
	}

	// pp = append(pp, paths[index])
	return MesingPath
}

func Small(ScorofPathe*[]int) int {
	var minscor int
	var index int
	minscor = (*ScorofPathe)[0]
	for r, v := range *ScorofPathe {
		if minscor == -1 && v != -1 {
			minscor = v
			index = r
		}
		if v == -1 {
			continue
		}
		if v < minscor {
			minscor = v
			index = r
		}
	}
	(*ScorofPathe)[index] = -1
	// fmt.Println(*p)
	return index
}

func check(path []string,CheckVisited *map[string]bool) bool {
	temp := []string{}
	for _, room := range path {
		if !(*CheckVisited)[room] {
			(*CheckVisited)[room] = true
			temp = append(temp, room)
		} else {
			for _, v := range temp {
				delete((*CheckVisited), v)
			}
			// fmt.Println("false")
			return false
		}
	}
	// fmt.Println("true")
	return true
}
