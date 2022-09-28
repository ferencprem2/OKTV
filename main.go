package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Enter a number: ")
	var number int
	fmt.Scanln(&number)
	divide(number, nil)
}

func divide(number int, currentNumber []int) {
	if 1 <= number && number <= 1000000000 {
		toString := strconv.Itoa(number)
		stringDividend := toString[0 : len(toString)-1]
		stringDivisor := toString[len(toString)-1:]
		dividend, _ := strconv.Atoi(stringDividend)
		divisor, _ := strconv.Atoi(stringDivisor)
		counter := 0

		numbers := currentNumber
		if (dividend-(11*divisor))%37 == 0 {
			fmt.Println("IGEN")
			for i := 0; i < len(numbers); i++ {
				fmt.Println(numbers[i])
			}
		} else if (dividend-(11*divisor))%37 != 0 {
			numbers = append(numbers, dividend-(11*divisor))
			counter++
			divide(numbers[counter], numbers)
		}
	}
}
