package Lem

func Chekslayce(Path []string, s string) bool { ///nchof wax xi eliment m3awd
	for _, v := range Path {
		if v == s {
			return true
		}
	}
	return false
}
