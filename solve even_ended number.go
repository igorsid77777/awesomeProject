// берем все пары 4-х значных чисел и пытаемся сосчитать соличество тех
// у которых первая и последняя цифра равны
// например 2804 и 2994

package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	counter := 0
	s := ""
	t := time.Now()
	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ {
			//s := fmt.Sprintf("%d", a * b)
			s = strconv.FormatInt(int64(a*b), 10)
			//String(a*b)
			if s[0] == s[len(s)-1] {
				counter += 1
			}
		}
	}
	fmt.Println(time.Since(t))
	fmt.Println(counter)
	counter = 0
	t = time.Now()
	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ {
			s = fmt.Sprintf("%d", a*b)
			if s[0] == s[len(s)-1] {
				counter += 1
			}
		}
	}
	fmt.Println(time.Since(t))
	fmt.Println(counter)

	for a := 10; a <= 99; a++ {
		for b := a; b <= 99; b++ {
			fmt.Printf("%d  -  %d \n", a, b)
		}
	}
}
