package main

import (
	"fmt"
	"log"
)

func main() {
	var err error

	r, err := NewClient("localhost:6379", 0)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// set
	err = r.Set(ctx, "key", "value")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("set completed")

	// get
	val, err := r.Get(ctx, "key")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("val -> ", val)

	// del
	err = r.Del(ctx, "key")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("del completed")

	// get
	val, err = r.Get(ctx, "key")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("val -> ", val)
}
