package common

import "fmt"

func Check(e error) {
	if e != nil {
		fmt.Println(e.Error())
	}
}