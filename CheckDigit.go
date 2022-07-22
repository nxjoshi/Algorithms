package main

import "fmt"

func main() {
	var num int
	var digit = 629104150021
	num = digit
	var arr [13]int
	var total int

	for i := 1; i <= 12; i++ {
		if num > 0 {
			digit = num % 10
			arr[i] = digit
			num = num / 10
			fmt.Println("digit:", arr[i])
		}

		if i%2 == 0 {
			arr[i] = arr[i] * 1
		} else {
			arr[i] = arr[i] * 3
		}
		fmt.Println("New digit:", arr[i])
		fmt.Println()
		total = total + arr[i]

	}

	fmt.Println("Total is : ", total)
	if total > 0 && total < 10 {
		total = total - 10
	}
	if total > 10 && total < 20 {
		total = 20 - total
	}
	if total > 20 && total < 30 {
		total = 30 - total
	}
	if total > 30 && total < 40 {
		total = 40 - total
	}
	if total > 40 && total < 50 {
		total = 50 - total
	}
	if total > 50 && total < 60 {
		total = 60 - total
	}

	if total > 60 && total < 70 {
		total = 70 - total
	}
	if total > 70 && total < 80 {
		total = 80 - total
	}

	if total > 80 && total < 90 {
		total = 90 - total
	}

	if total > 90 && total < 110 {
		total = 100 - total
	}
	fmt.Println("Result is : ", total)
}
