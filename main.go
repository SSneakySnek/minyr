package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = scanner.Text()
		if input == "q" || input == "exit" {
			fmt.Println("Exit")
			os.Exit(0)
		} else if input == "convert" {
			fmt.Println("Konverterer alle målingene gitt i grader Celsius til grader Farenheit.")

		} else {
			fmt.Println("Venligst velg convert, average eller exit:")
		}
	}
	if err != nil {
		log.Fatal(err)
	}
}
