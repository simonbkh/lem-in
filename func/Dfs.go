package Lem

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
