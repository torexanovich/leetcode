package main

// import (
// 	"fmt"
// )

// func main() {
// 	candies := []int{12, 1, 12}
// 	fmt.Println(kidsWithCandies(candies, 10))
// }

// Use greedy approach. For each kid check if candies[i] + extraCandies ≥ maximum in Candies[i].

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max, ans := candies[0], make([]bool, len(candies))

	for i, v := range candies {
		if v > max {
			max = v
		}

		ans[i] = v+extraCandies >= max
	}

	return ans 
}
