package shell

import "mygoshell.com/s/shellerr"

func M(strs []string) bool {
	script := setScript()
	scriptNom := isScript(script, strs[0])
	if scriptNom == -1 {
		shellerr.Nomal("command notfound.")
		return false
	}
	return true
}

func isScript(s []string, a string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == a {
			return i
		}
	}
	return -1
}
