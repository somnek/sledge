package main

import (
	"context"
)

func main() {
	ctx := context.Background()
	dummyIns()
	Ping(ctx)
	show(ctx)
}
