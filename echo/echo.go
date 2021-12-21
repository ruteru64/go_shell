package echo

import (
	"fmt"
	"os"

	"mygoshell.com/s/shellerr"
)

func Script() bool {
	if len(os.Args) == 2 {
		return false
	}
	if len(os.Args) == 3 {
		fmt.Println(os.Args[2])
	}
	if len(os.Args) == 4 {
		shellerr.Nomal("echo parms err")
		return false
	}
	if len(os.Args) == 5 {
		if os.Args[3] == "->" {
			f, err := os.OpenFile(os.Args[4], os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				shellerr.Nomal("fail to open file")
			}
			defer f.Close()
			n, err := f.WriteString(os.Args[2])
			if err != nil {
				shellerr.Nomal("fail to write file")
			}
			return n >= 1
		} else {
			shellerr.Nomal("echo parms err")
			return false
		}
	}
	return true
}
