package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

var format string

func main() {
	// Read in command line argument
	str := os.Args[1]

	base := uint(0)
	powers := []uint{}
	maxPower := uint(len(str) - 1)

	// Here, we read though the string, treat it as a binary number and get:
	// 1. The base value, replacing all X's with 0
	// 2. All X values as a number if all other numbers were 0 (powers)
	for i, r := range str {
		power := maxPower - uint(i)
		switch r {
		case 'X':
			powers = append(powers, 1<<power)
		case '1':
			base += 1 << power
		case '0':
		default:
			fmt.Println("Invalid input")
			os.Exit(1)
		}
	}

	// We have a dynamically sized format to pad 0's in front of the number
	format = "%0" + fmt.Sprintf("%d", utf8.RuneCountInString(str)) + "b\n"

	printVal(base)
	calculate(base, 0, powers)
}

// Recursively calculate by adding all combinations of the powers
func calculate(val uint, powerIndex int, powers []uint) {
	valPlus := val + powers[powerIndex]

	printVal(valPlus)
	if powerIndex < len(powers)-1 {
		calculate(val, powerIndex+1, powers)
		calculate(valPlus, powerIndex+1, powers)
	}
}

// Convert the number back into binary format, and print
func printVal(val uint) {
	fmt.Printf(format, val)
}
