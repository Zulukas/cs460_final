package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Location struct {
	x int8
	y int8
}

func display(loc Location) {
	fmt.Print("(" + strconv.Itoa(int(loc.x)) + ", " + strconv.Itoa(int(loc.y)) + ")")
}

func isValid(loc Location) bool {
	return (loc.x >= 0 && loc.x <= 7 && loc.y >= 0 && loc.y <= 7)
}

func getDelta(src Location, dst Location) (int8, int8) {
	return src.x - dst.x, -1 * (src.y - dst.y)
}

func getLocationBetween(src Location, dst Location) (Location, bool) {
	xDelta, yDelta := getDelta(src, dst)

	if abs(xDelta) != 2 || abs(yDelta) != 2 {
		return Location{-1, -1}, false //Invalid location!
	}

	loc := Location{0, 0}

	if dst.x > src.x {
		loc.x = dst.x - 1
	} else {
		loc.x = src.x - 1
	}

	if dst.y > src.y {
		loc.y = dst.y - 1
	} else {
		loc.y = src.y - 1
	}

	return loc, true
}

func interpretLocation(input string) Location {
	runes := []rune(strings.ToUpper(input))

	x_coordinate := int8(runes[0]) - 65
	y_coordinate := int8(runes[1]) - 49

	if x_coordinate < 0 || x_coordinate > 7 || y_coordinate < 0 || y_coordinate > 7 {
		x_coordinate = -1
		y_coordinate = -1
	}

	loc := Location{x_coordinate, y_coordinate}

	return loc
}