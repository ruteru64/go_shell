package pwd

import (
	"fmt"
	"os"
)

func Script() bool {
	p, _ := os.Getwd()
	fmt.Println(p)
	return true
}
