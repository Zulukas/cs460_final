package main

func abs(val int8) int8 {
	if val < 0 {
		return val * -1
	} else {
		return val
	}
}