package Lem


func Find(sl [][]string, s string) int {////had func bax nl9a nmla axmn Path fin kayna 
	for x, v := range sl {
		for _, va := range v {
			if va == s {
				return x
			}
		}
	}
	return 0
}