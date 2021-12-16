package main

import (
	"os"

	"mygoshell.com/s/shell"
	"mygoshell.com/s/shellerr"
)

func main() {
	if len(os.Args) == 1 {
		shellerr.Nomal("Please add command.")
		return
	}
	err := shell.M(os.Args)
	if !err {
		return
	}
}
