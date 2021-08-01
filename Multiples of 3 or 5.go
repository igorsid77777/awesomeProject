/*
	нужно вывести сумму чисел от 1 до N кратных 3 или 5  которые меньше введеного числа N
*/

package main

import (
	"fmt"
)

func Multiple3And5(number int) int {
	/*	i := 3
		numForSum := make([]int, 0)
		for i < number {

			if i % 3 == 0 && i % 5 == 0 {
				numForSum = append(numForSum, i)
			}else if i % 3 == 0 {
				numForSum = append(numForSum, i)
			}else if i % 5 == 0{
				numForSum = append(numForSum, i)
			}
			i++
		}
		k := 0
		for j:= range numForSum{
			k += numForSum[j]
		}
		return k*/

	i := 3
	numForSum := 0
	for i < number {

		if i%3 == 0 && i%5 == 0 {
			numForSum += i
		} else if i%3 == 0 {
			numForSum += i
		} else if i%5 == 0 {
			numForSum += i
		}
		i++
	}
	return numForSum
}

func main() {
	fmt.Println(Multiple3And5(15))
}
