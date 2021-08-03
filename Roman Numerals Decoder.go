/*
Create a function that takes a Roman numeral as its argument and returns its value
as a numeric decimal integer. You don't need to validate the form of the Roman numeral.
декодер римских чисел
*/
package main

import "fmt"

func Decode(roman string) int {
	var RomeValue = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	value := 0
	currentValue := 0
	lastValue := 0
	for i := len(roman) - 1; i >= 0; i-- {
		currentValue = RomeValue[roman[i]]
		if currentValue < lastValue {
			value -= currentValue
		} else {
			value += currentValue
		}
		lastValue = currentValue
	}
	return value
}

func main() {
	fmt.Println(Decode("CMV"))
	fmt.Println(Decode("MMDCCXVI"))
	fmt.Println(Decode("MXCIX"))
}
