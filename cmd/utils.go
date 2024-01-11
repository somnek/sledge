package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
)

func sortFVPair(fv *[]FVPair) {
	// todo
	sort.Slice(*fv, func(i, j int) bool {
		return (*fv)[i].field < (*fv)[j].field
	})
}

func countRune(s string, r rune) int {
	var count int
	for _, c := range s {
		if c == r {
			count++
		}
	}
	return count
}

func logToFile(s ...string) {
	f, err := os.OpenFile("debug.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o600)
	if err != nil {
		fmt.Println("error opening file for logging:", err)
		os.Exit(1)
	}
	defer f.Close()

	// write to file
	log.SetOutput(f)
	for _, e := range s {
		log.Println(e)
	}

}
