package make

import (
	"fmt"
	"io/fs"
	"os"
	"strconv"
)

func Touch() bool {
	_, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func Mkdir() bool {
	var err error
	if len(os.Args) < 3 {
		fmt.Println("add filename")
		return false
	} else if len(os.Args) == 4 {
		i, _ := strconv.Atoi(os.Args[3])
		u := uint32(i)
		a := fs.FileMode(u)
		err = os.Mkdir(os.Args[2], a)
	} else {
		err = os.Mkdir(os.Args[2], 0777)
	}
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
