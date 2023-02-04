package main

import "context"

var dummyData = map[string]string{
	"a": "🍎",
	"b": "🍌",
	"c": "🐈",
	"d": "🐕",
	"e": "🐘",
}

func dummyIns(ctx context.Context) {
	for key, val := range dummyData {
		add(ctx, key, val)
	}
}
