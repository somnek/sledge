package main

func contains[T int | string](slice []T, item T) bool {
	for _, x := range slice {
		if x == item {
			return true
		}
	}
	return false
}

func remove[T int | string](slice []T, item T) []T {
	var newSlice []T
	for _, x := range slice {
		if x != item {
			newSlice = append(newSlice, x)
		}
	}
	return newSlice
}
