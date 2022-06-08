package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[0] = "a"
	m[1] = "b"
	for k, v := range m {
		fmt.Printf("%d | %s\n", k, v)
	}
}
