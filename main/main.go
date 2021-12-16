package main

import (
	"os"

	"../shellerr"
)

func main() {
	if len(os.Args) == 1 {
		shellerr.Nomal("Please add command.")
		return
	}
}
