package main

import "fmt"

func main() {
	str := "Helmi Yusuf Efendi     "
	slice := make([]string, 0)
	for i := len(str) - 1; i >= 0; i-- {
		if string(str[i]) == string(" ") {
			continue
		} else {
			slice = append(slice, string(str[i]))
		}
	}
	fmt.Println(slice)
}