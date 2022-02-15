package main

import "fmt"

/* You are making a robot that can speak numbers.
Your robot should take 3 numbers in the range of 0-10 as input and output the corresponding texts in English.

Sample Input:
8
0
5

Sample Output:
Eight
Zero
Five */

func main() {
	fmt.Println("Insert a number from 0 to 10 to get its value in letters")

	var firstNumber, secondNumber, thirdNumber int
	fmt.Print("Insert first number: ")
	numberInput(&firstNumber)
	fmt.Print("Insert second number: ")
	numberInput(&secondNumber)
	fmt.Print("Insert third number: ")
	numberInput(&thirdNumber)

	fmt.Println("\nNumbers converted:")
	printConversion(firstNumber)
	printConversion(secondNumber)
	printConversion(thirdNumber)
}

func printConversion(number int) {
	fmt.Printf("%d -> %s\n", number, numberToWord(number))
}

func numberInput(number *int) {
	fmt.Scanln(number)
	if *number > 10 || *number < 0 {
		fmt.Println("Not valid. Insert a number from 0 to 10")
		numberInput(number)
	}
}

func numberToWord(number int) string {
	switch number {
	case 0:
		return "Zero"
	case 1:
		return "One"
	case 2:
		return "Two"
	case 3:
		return "Three"
	case 4:
		return "Four"
	case 5:
		return "Five"
	case 6:
		return "Six"
	case 7:
		return "Seven"
	case 8:
		return "Eight"
	case 9:
		return "Nine"
	case 10:
		return "Ten"
	default:
		return "Not a valid number"
	}
}
