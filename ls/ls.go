package ls

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Script() bool {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println(err)
		return false
	}
	isa := false
	if len(os.Args) > 2 {
		if os.Args[2] == "-a" {
			isa = true
		}
	}
	for _, f := range files {
		if !isa && f.Name()[:1] == "." {
			continue
		}
		if f.IsDir() {
			fmt.Println("\x1b[32m" + f.Name() + "/" + "\x1b[0m")
		} else {
			fmt.Println(f.Name())
		}
	}
	return true
}
