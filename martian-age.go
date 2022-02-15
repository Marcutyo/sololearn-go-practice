package main

import "fmt"

func main() {
	var earthAge int
	fmt.Print("🌍 Insert your age on Earth: ")
	fmt.Scanln(&earthAge)
	fmt.Println(earthAge)

	marsAge := earthToMars(earthAge)
	fmt.Printf("👽 Your Martian age is %v", marsAge)
}

func earthToMars(earthAge int) int {
	days := yearToDays(earthAge)
	return marsAge(days)
}

func yearToDays(year int) int {
	return year * 365
}

func marsAge(days int) int {
	return days / 687
}
