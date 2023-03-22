package yr

import (
	"fmt"
	"path/to/yr"
)

func test() {
	filepath := "C:/Users/thomas/Desktop/visualCoding/minyr/kjevik-temp-celsius-20220318-20230318.csv"
	numLines, err := yr.CountLinesInCSVFile(filepath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Number of lines in the CSV file: %d", numLines)
}
