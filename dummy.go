package main

var data = map[string]string{
	"h": "place holder...",
	"a": "apple",
	"b": "two dog in da house, two balls in my mouth",
	"c": "bake a cake",
	"d": "true",
}

func dummyIns() {
	rdb := connect(0)
	for key, value := range data {
		rdb.add(key, value)
	}
}
