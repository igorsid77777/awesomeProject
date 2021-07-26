package main

import (
	"fmt"
	str "strings"
)

func ToCamelCase(s string) string {
	var sArr []string = str.Split(s, "-")
	if len(sArr) == 1 {
		sArr = str.Split(s, "_")
	}

	for i := 1; i < len(sArr); i++ {
		if str.ToUpper(string(sArr[i][0])) != string(sArr[i][0]) {
			sArr[i] = str.Replace(sArr[i], string(sArr[i][0]), str.ToUpper(string(sArr[i][0])), 1)
		}
	}
	return str.Join(sArr, "")
}

// Function foo accepts a slice of integers and returns the slice back.
func foo(slice []int) []int {
	slice[0] = 10
	slice[1] = 20
	return slice
}

func main() {
	st := "the-stealth-warrior"
	fmt.Println(ToCamelCase(st))
	st = "The_Stealth_Warrior"
	fmt.Println(ToCamelCase(st))
	var S string
	S = "qwerty"
	fmt.Println(string(S[0:3]))
	fmt.Printf("%c %s\n", S[5], S[1:3])
	fmt.Println("**************************************")
	question := "¿Cómo estás?"

	for _, c := range question {
		fmt.Printf("%c", c)
	}
	fmt.Println()
	slice := make([]int, 10)

	slice[0] = 100
	slice[1] = 200
	fmt.Println(slice)
	slice = foo(slice)
	fmt.Println(slice)
}
