package main

import "fmt"

func main() {
	arr := [15]int{3, 7, 8, 2, 8, 2, 2, 4, 6, 3, 1, 0, 0, 0, 5}
	fmt.Println(luhn(arr))
}

func luhn(a [15]int) bool {
	checksum := a[14]
	sum := 0
	for i := 13; i >= 0; i-- {
		digit := a[i]
		if (13-i)%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	sum += checksum
	if sum%10 == 0 {
		return true
	}
	return false
}
