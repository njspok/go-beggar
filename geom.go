package main

import "math"

func Distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
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
