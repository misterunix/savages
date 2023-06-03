package main

import "math"

// Returns the X and Y of the index.
func Index2XY(i int) (int, int) {
	return i % maxX, i / maxX
}

// Returns the index of the X and Y.
func XY2Index(x, y int) int {
	return y*maxX + x
}

// Returns the distance between two XY points.
func Distance(x1, y1, x2, y2 int) int {
	return int(math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))))
}

// Returns the distance between two index points.
func DistanceIndex(i1, i2 int) int {
	x1, y1 := Index2XY(i1)
	x2, y2 := Index2XY(i2)
	return Distance(x1, y1, x2, y2)
}

// Returns the distance between two savages.
func DistanceSavage(s1, s2 savage) int {
	return DistanceIndex(s1.Location, s2.Location)
}
