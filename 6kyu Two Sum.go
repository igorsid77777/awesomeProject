/*
Write a function that takes an array of numbers (integers for the tests) and a target number.
It should find two different items in the array that, when added together,
give the target value.
The indices of these items should then be returned in a tuple like so: (index1, index2).
*/

package main

import "fmt"

func TwoSum(numbers []int, target int) [2]int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			sum := numbers[i] + numbers[j]
			//fmt.Printf("       %d  %d\n", numbers[i],numbers[j])
			if sum == target {
				//fmt.Printf("%d  %d\n", numbers[i],numbers[j])
				return [2]int{i, j}
			}
		}
	}
	return [2]int{0, 0}
}

func main() {
	fmt.Println(TwoSum([]int{1, 2, 3, 0, 4, 2}, 4))
	fmt.Println(TwoSum([]int{1234, 5678, 9012}, 14690))
	fmt.Println(TwoSum([]int{2, 2, 3}, 4))
}
