package main

import (
	"context"
	"testing"
)

func TestAdd(t *testing.T) {
	ctx := context.Background()
	tKey := "primeagen"
	tVal := "rust"
	add(ctx, tKey, tVal)

	r := get(ctx, tKey)
	if r != tVal {
		t.Errorf("add() failed. Expected %s, got %s", tKey, tVal)
	} else {
		t.Log("pased")
	}
}
