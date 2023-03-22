package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("minyr/kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Det er så mange linjer i filen: %d", len(lines))
}
