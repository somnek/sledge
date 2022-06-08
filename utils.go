package main

import "strings"

func glueLogs(logs *[]string) (s string) {
	return strings.Join(*logs, "\n")
}

func valType(val string) (s string) {
	for _, c := range val {
		if string(c) == "{" {
			return "map"
		}
	}
	return "str"
}
