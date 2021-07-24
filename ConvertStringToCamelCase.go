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

func main() {
	st := "the-stealth-warrior"
	fmt.Println(ToCamelCase(st))
	st = "The_Stealth_Warrior"
	fmt.Println(ToCamelCase(st))
}
