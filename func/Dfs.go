package Lem

import "slices"

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
		ind := -1
		if start != m.start {
			ind = slices.Index(Link[start], end)
			if ind != -1 {
				currentPath = append(currentPath, Link[start][ind])
				*paths = append(*paths, append([]string{}, currentPath...))
			}
		}

		if ind == -1 {
			for _, neighbor := range Link[start] {
				if !visited[neighbor] {
					dfs(neighbor, end, visited, currentPath, paths, m)
				}
			}
		}
	}

	visited[start] = false
}
