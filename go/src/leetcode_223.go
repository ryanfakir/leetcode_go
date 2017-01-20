package main

import "math"

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	left := math.Max(float64(A), float64(E))
	right := math.Max(left, math.Min(float64(C), float64(G)))
	bottom := math.Max(float64(B), float64(F))
	top := math.Max(bottom, math.Min(float64(D), float64(H)))
	return (D-B)*(C-A) + (H-F)*(G-E) - int(top-bottom)*int(right-left)
}
