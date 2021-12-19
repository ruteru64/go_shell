package shell

import (
	"mygoshell.com/s/a"
	"mygoshell.com/s/cat"
	"mygoshell.com/s/ls"
	"mygoshell.com/s/pwd"
	"mygoshell.com/s/shellerr"
)

func M(strs []string) bool {
	if isFlag(strs[1]) {
		return flag(strs)
	} else {
		return script(strs)
	}
}

func flag(sarry []string) bool {
	flags := setflag()
	flagNom := isListFlag(flags, sarry[1])
	if flagNom == -1 {
		shellerr.Nomal("flag notfound.")
		return false
	}
	return true
}

func script(sarry []string) bool {
	scripts := setScript()
	scriptNom := isScript(scripts, sarry[1])
	if scriptNom == -1 {
		shellerr.Nomal("command notfound.")
		return false
	}
	if scriptNom == 0 {
		a.Script()
		return true
	} else if scriptNom == 1 {
		return pwd.Script()
	} else if scriptNom == 2 {
		return cat.Script()
	} else if scriptNom == 3 {
		return ls.Script()
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

func isFlag(s string) bool {
	if s[0] == '-' {
		return true
	}
	return false
}

func isListFlag(s []string, a string) int {
	a = a[1:]
	for i := 0; i < len(s); i++ {
		if s[i] == a {
			return i
		}
	}
	return -1
}
