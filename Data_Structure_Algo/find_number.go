package datastructurealgo

import "fmt"

func hasEvenDigit(n int) bool {
	digitCount := 0
	for n > 0 {
		digitCount++
		n /= 10
	}
	return digitCount%2 == 0
}

func findNumbers(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if hasEvenDigit(nums[i]) {
			count++
		}
	}
	return count
}

func RunFindNumber() {
	fmt.Println("Total count of numbers having even number of digits in array")
	nums := [5]int{12, 345, 2, 6, 7896}
	result := findNumbers(nums[:])
	fmt.Printf("Count of numbers having even number of digits: %d\n", result)
}
