package tools

import "fmt"

func ReStart() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
