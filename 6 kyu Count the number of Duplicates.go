/*
Write a function that will return the count of distinct case-insensitive alphabetic characters
and numeric digits that occur more than once in the input string.
The input string can be assumed to contain only alphabets (both uppercase and lowercase) and numeric digits.
*/

package main

import (
	"fmt"
	"strings"
)

/*
func duplicate_count(s1 string) int {
	s1 = strings.ToUpper(s1)
	fmt.Println(s1)
	counter := 0
	check := make(map[string]int)
	for i := 0; i < len(s1); i++ {
		//fmt.Printf("%T", s1[i])
		s11 := string(s1[i])
		if _, ok := check[fmt.Sprintf("%s", s11)]; !ok{
			count := strings.Count(s1, fmt.Sprintf("%s", s11))
			if count > 1 {
				counter += 1
			}
			check[fmt.Sprintf("%s", s11)] = count
		}
	}
	return counter
}
*/
/*
func duplicate_count(s1 string) (counter int) {
	s1 = strings.ToUpper(s1)
	counter = 0
	for i := 0; i < len(s1); i++ {
		//fmt.Printf("%s\n", s1)
		s11 := string(s1[i])
		if strings.Count(s1, fmt.Sprintf("%s", s11)) > 1 {
				counter += 1
		}
		s1 = strings.ReplaceAll(s1, string(s11), "")
	}
	return counter
}
*/

func duplicate_count(s string) (r int) {
	t := map[rune]int{}
	for _, c := range strings.ToLower(s) {

		//t[c] ++
		t[c] = t[c] + 1
	}
	for _, n := range t {
		if n > 1 {
			r++
		}
	}
	return
}

func main() {
	s := "Hello"
	t := map[rune]int{}
	t[rune(s[0])] = 1
	fmt.Println(rune(s[1]))
	fmt.Println(string(72), string(101))

	fmt.Println(duplicate_count("aA11"))
	fmt.Println(duplicate_count("abcdea"))
	fmt.Println(duplicate_count("abcdeaB11"))
	fmt.Println(duplicate_count("indivisibility"))
}
