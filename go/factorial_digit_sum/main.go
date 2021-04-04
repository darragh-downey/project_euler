package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	num := int64(100)
	s := factoSum(num)
	fmt.Printf("%d", s)
}

func factoSum(num int64) int64 {
	result := big.NewInt(0)
	result.MulRange(1, num)
	c := result.String()

	sum := int64(0)
	for idx, i := range c {
		is := string(i)
		fmt.Printf("%d - %s\n", idx, is)
		s, err := strconv.ParseInt(is, 10, 64)
		if err != nil {
			fmt.Printf("E: %v", err)
		}
		sum += s
	}
	return sum
}
