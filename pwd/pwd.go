package pwd

import (
	"fmt"
	"os"
)

func Pwd() string {
	p, _ := os.Getwd()
	return p
}

func Script() bool {
	p, _ := os.Getwd()
	fmt.Println(p)
	return true
}
