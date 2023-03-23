package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/SSneakySnek/minyr/yr"
)

func main() {
	for {
		fmt.Println("Velg convert, average eller exit:")
		input := readInput()
		switch input {

		case "convert":
			fmt.Println("Konverterer alle målingene i grader Celsius til grader Fahrenheit.")
			yr.ConvertTemperature()

		case "average":
			fmt.Println("Gjennomsnitt-kalkulator")
			yr.AverageTemperature()

			for {
				fmt.Println("Avslutte? (y/n)")
				input2 := readInput()
				if input2 == "y" {
					break
				} else if input2 == "n" {
					yr.AverageTemperature()
				} else {
					fmt.Println("Ugyldig input, prøv igjen")
				}
			}

		case "exit":
			fmt.Println("Programmet avsluttes.")
			return
		default:
			fmt.Println("Ugyldig input, prøv igjen")
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
