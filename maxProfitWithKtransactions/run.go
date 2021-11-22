package main

import "fmt"

func main() {
	prices := []int{5, 11, 3, 50, 60, 90}
	k := 2
	fmt.Println(MaxProfitWithKTransactions(prices, k))
}

func MaxProfitWithKTransactions(prices []int, k int) int {
	if len(prices) == 0 {
		return 0
	}
	prices_matrix := make([][]int, k+1, len(prices))
	for i := range prices_matrix {
		prices_matrix[i] = make([]int, len(prices))
	}
	fmt.Printf("%v", prices_matrix)
	return -1
}
