package main 

import "fmt"


func main(){
	m := make(map[int]string)
	m[0] = "zero"
	m[1] = "one"

	for k, v := range m {
		fmt.Println(k, v)
	}
}

