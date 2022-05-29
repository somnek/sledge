package main

import "strings"

func glueLogs(logs *[]string) (s string) {
	return strings.Join(*logs, "\n")
}
