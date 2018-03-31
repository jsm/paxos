package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

var format string

func main() {
	str := os.Args[1]

	base := uint(0)
	powers := []uint{}
	maxPower := uint(len(str) - 1)

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
	format = "%0" + fmt.Sprintf("%d", utf8.RuneCountInString(str)) + "b\n"

	printVal(base)
	calculate(base, 0, powers)
}

func calculate(val uint, powerIndex int, powers []uint) {
	valPlus := val + powers[powerIndex]

	printVal(valPlus)
	if powerIndex < len(powers)-1 {
		calculate(val, powerIndex+1, powers)
		calculate(valPlus, powerIndex+1, powers)
	}
}

func printVal(val uint) {
	fmt.Printf(format, val)
}
