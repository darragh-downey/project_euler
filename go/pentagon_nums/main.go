package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// find the smallest pair of pentagonal numbers
	// where their sum and difference are both pentagonal
	start1 := time.Now()
	pents := make([]int, 0)
	for i := 1; i < 10000; i++ {
		pents = append(pents, pen(i))
	}

	k, j := bruteLocate(pents)

	d := math.Abs(float64(k) - float64(j))
	end1 := time.Since(start1)
	fmt.Printf("%f Completed in %s\n", d, end1)

	start2 := time.Now()
	d2 := smarterWay()
	end2 := time.Since(start2)
	fmt.Printf("%d Completed in %s\n", d2, end2)
}

func bruteLocate(pents []int) (int, int) {
	for i := 0; i < len(pents)-1; i++ {
		for j := 1; j < len(pents)-1; j++ {
			pentSum := pents[i] + pents[j]
			pentDiff := pents[j] - pents[i]
			foundSum := false
			foundDiff := false

			for _, p := range pents {
				if p == pentSum {
					foundSum = true
				}
				if p == pentDiff {
					foundDiff = true
				}

				if foundDiff && foundSum {
					fmt.Printf("Found! %d %d\n", pentSum, pentDiff)
					break
				}
			}

			if foundSum && foundDiff {
				fmt.Printf("%d %d\n", pents[i], pents[j])
				return pents[i], pents[j]
			}
		}
	}
	// not found
	return -1, -1
}

func smarterWay() int {
	// https://blog.dreamshire.com/project-euler-44-solution/
	ps := Set{}
	i := 1000
	for {
		i += 1
		s := (3*i*i - i) / 2
		for Pj, _ := range ps.items {
			if ps.Has(s-Pj) && ps.Has(s-2*Pj) {
				return s - 2*Pj
			}
		}
		ps.Add(s)
	}
}

func pen(num int) int {
	return num * (3*num - 1) / 2
}

type Set struct {
	items map[int]bool
}

func (s *Set) Add(i int) *Set {
	if s.items == nil {
		s.items = make(map[int]bool)
	}

	_, ok := s.items[i]
	if !ok {
		s.items[i] = true
	}
	return s
}

func (s *Set) Clear() {
	s.items = make(map[int]bool)
}

func (s *Set) Delete(i int) bool {
	_, ok := s.items[i]
	if ok {
		delete(s.items, i)
	}
	return ok
}

func (s *Set) Has(i int) bool {
	_, ok := s.items[i]
	return ok
}

func (s *Set) Items() []int {
	items := []int{}
	for i := range s.items {
		items = append(items, i)
	}
	return items
}

func (s *Set) Size() int {
	return len(s.items)
}
