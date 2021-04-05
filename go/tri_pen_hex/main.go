package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Printf("%d", clever())
}

func clever() int {
	p := 165
	h := 143
	h = 84*p + 97*h - 38
	return h * (2*h - 1)
}

func brute(upper int) {
	for i := 1; i <= upper; i++ {
		t := tri(i)
		for j := 1; j <= upper; j++ {
			p := pen(j)
			for k := 1; k <= upper; k++ {
				h := hexa(k)
				if t == p && p == h {
					log.Println(i, j, k)
					break
				}
			}
		}
	}
}

func tri(num int) int {
	return num * (num + 1) / 2
}

func pen(num int) int {
	return num * (3*num - 1) / 2
}

func penTri(num int) int {
	return tri(num) + 2*tri(num-1)
}

func hexa(num int) int {
	return num * (2*num - 1)
}

func hexaTri(num int) int {
	return 6*tri(num-1) + 1
}

// Pentagonal Triangular Number
// https://mathworld.wolfram.com/PentagonalTriangularNumber.html
// 1/2 n(3n - 1) = 1/2 m(m+1)
// (6n - 1)^2 - 3(2m + 1)^2 = -2
// Substiuting x = 6n-1 and y = 2m+1
// x^2 - 3y^2 = -2
// or
// 36n^2 - 48m^2 - 12n + 24m
// feed into our favourite Diophantine equation solver and produce the results:
// P0 = 0
// H0 = 0
//
// Pn+1 = 97 * Pn + 112 * Hn - 44
// Hn+1 = 84 * Pn + 97 * Hn - 38

// Replace P0 and H0 with the starting index and this problem is solved instantly.
// NOTE: the starting indexes were defined in the problem.
