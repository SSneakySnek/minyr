package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/SSneakySnek/minyr/yr"
)

func main() {
	for {
		fmt.Println("Choose convert, average or exit:")
		input := readInput()
		switch input {

		case "convert":
			fmt.Println("Converting all units of Celsius to Farenheit.")
			yr.ConvertTemperature()

		case "average":
			fmt.Println("Average temperature calculator")
			yr.AverageTemperature()

			readInput() // Add this line to clear any extra newline characters from the input buffer

			for {
				fmt.Println("Quit? (y/n)")
				input2 := readInput()
				if input2 == "y" {
					break
				} else if input2 == "n" {
					yr.AverageTemperature()
				} else {
					fmt.Println("Invalid input, try again")
				}
			}

		case "exit":
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid input, try again")
		}
	}
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
}
