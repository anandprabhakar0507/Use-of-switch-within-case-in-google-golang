package main

import "fmt"

func main() {
	for day := 1; day <= 12; day++ {
		fmt.Println("on the", day, "of christmas my true love sent to me")

		switch day {
		case 12:
			fmt.Println("twelve drummers drumming")
		case 11:
			fmt.Println("Eleven pipers piping")
		case 10:
			fmt.Println("ten lords a leaping")
		case 9:
			fmt.Println("nine ladies dancing")
		case 8:
			fmt.Println("eight maids a milking")
		case 7:
			fmt.Println("seven swans swimming")
		case 6:
			fmt.Println("six geese a laying")
		case 5:
			fmt.Println("five golden rings")
		case 4:
			fmt.Println("four calling birds")
		case 3:
			fmt.Println("three french hens")
		case 2:
			fmt.Println("two turtle doves")
		case 1:
			fmt.Println("a partridge in fear tree.")
		}

	}
}
