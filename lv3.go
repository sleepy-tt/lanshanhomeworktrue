package main

import "fmt"

func assess(number int) bool {
	if number < 2 {
		return false
	}
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	number := 13
	if assess(number) {
		fmt.Println(number, "是素数")
	} else {
		fmt.Println(number, "不是素数")
	}
}
