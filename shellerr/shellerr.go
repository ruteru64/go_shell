package shellerr

import "fmt"

func Nomal(m string) bool {
	fmt.Printf("\x1b[31m%s\x1b[0m\n", m)
	return true
}
