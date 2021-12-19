package cat

import (
	"fmt"
	"os"

	"mygoshell.com/s/pwd"
)

func Script() bool {
	p := pwd.Pwd() + "\\" + os.Args[2]
	fmt.Println(p)
	f, err := os.Open(p)
	if err != nil {
		fmt.Println(err)
		return false
	}
	size := 1024 * 1024
	data := make([]byte, size)
	count, err := f.Read(data)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(string(data[:count]))
	return true
}

func Cat() string {
	p := pwd.Pwd() + "\\" + os.Args[2]
	fmt.Println(p)
	f, err := os.Open(p)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	size := 1024 * 1024
	data := make([]byte, size)
	count, err := f.Read(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(data[:count])
}
