package yr

import (
	"bufio"
	"fmt"
	"os"
)

func testingForLines() {
	file, err := os.Open("thomas - Personal/Desktop/visualCoding/minyr/kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan() {
		lines++
	}

	fmt.Printf("Filen inneholder så mange linjer: %d", lines)
}
