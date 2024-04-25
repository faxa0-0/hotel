package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	var lastTwoDigits = num % 100

	switch {
	case lastTwoDigits == 11 || lastTwoDigits == 12 || lastTwoDigits == 13 || lastTwoDigits == 14:
		fmt.Printf("%d компьютеров", num)
	case num%10 == 1:
		fmt.Printf("%d компьютер", num)
	case num%10 >= 2 && num%10 <= 4:
		fmt.Printf("%d компьютера", num)
	default:
		fmt.Printf("%d компьютеров", num)
	}
}
