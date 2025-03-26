package main

import "math"

func Distance(start, end Point) float64 {
	return math.Sqrt(math.Pow(start.X-end.X, 2) + math.Pow(start.Y-end.Y, 2))
}

func IsCollision(x1, y1, w1, h1, x2, y2, w2, h2 float64) bool {
	if x1+w1 < x2 || x2+w2 < x1 {
		return false
	}
	if y1+h1 < y2 || y2+h2 < y1 {
		return false
	}
	return true
}
