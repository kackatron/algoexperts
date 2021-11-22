package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{5, 11, 3, 50, 60, 90}
	k := 2
	fmt.Println(MaxProfitWithKTransactions(prices, k))
}

func MaxProfitWithKTransactions(prices []int, k int) int {
	if len(prices) == 0 {
		return 0
	}

	profits := make([][]int, k+1)
	for i := range profits {
		profits[i] = make([]int, len(prices))
	}

	for tr := 1; tr < k+1; tr++ {
		fmt.Printf("%v\n\n", profits)
		max := math.MinInt64
		for day := 1; day < len(prices); day++ {
			fmt.Printf("%v\n", profits)
			max = int(math.Max(float64(max), float64(profits[tr-1][day-1]-prices[day-1])))
			profits[tr][day] = int(math.Max(float64(profits[tr][day-1]), float64(max+prices[day])))
		}
	}
	fmt.Printf("%v\n\n", profits)
	return profits[k][len(prices)-1]
}
