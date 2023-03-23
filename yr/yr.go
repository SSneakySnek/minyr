package yr

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/SSneakySnek/minyr/conv"
)

func ProcessLine(line string) string {
	fields := strings.Split(line, ";")
	if len(fields) != 4 {
		return line
	}
	temperature, err := strconv.ParseFloat(fields[3], 64)
	if err != nil {
		return line
	}
	fahrenheit := CelsiusToFahrenheit(temperature)
	fields[3] = strconv.FormatFloat(fahrenheit, 'f', 1, 64)
	fields[2] = "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Thomas"
	return strings.Join(fields, ";")
}

func CelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 1.8) + 32
}

func ConvertTemperature() {

	// Sjekker om filen allerede eksistere

	if _, err := os.Stat("output-test.csv"); err == nil {
		fmt.Print("Filen eksisterer allerede. Vil du generere filen på nytt? (y/n): ")
		var overwriteInput string

		fmt.Scanln(&overwriteInput)
		fmt.Println("Genererer filen på nytt...")

		if strings.ToLower(overwriteInput) == "n" {
			fmt.Println("Tilbake til hovedmeny")
			return
		}

	}

	// Åpner input fil
	file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Lager output fil
	outputFile, err := os.Create("output-test.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	// Lager en skriver for å skrive til output filen
	outputWriter := bufio.NewWriter(outputFile)

	// Lager en skanner for å lese input filen
	scanner := bufio.NewScanner(file)

	// Skriver første linjen til output filen
	if scanner.Scan() {
		_, err := outputWriter.WriteString(scanner.Text() + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	var outputLine string

	// Loop gjennom hver linje i input filen
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// Splitter linjene til sekjsoner
		fields := strings.Split(line, ";")
		var lastField string
		if len(fields) > 0 {
			lastField = fields[len(fields)-1]
		}

		// Konverterer siste field til farenheit hvis der er en
		var convertedField string
		if lastField != "" {
			var err error
			convertedField, err = convertLastField(lastField)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				continue
			}
		}

		// Bytter den orgianle siste field med en kovertert en
		if convertedField != "" {
			fields[len(fields)-1] = convertedField
		}
		if line[0:7] == "Data er" {
			outputLine = "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Thomas"
		} else {
			outputLine = strings.Join(fields, ";")

		}
		_, err = outputWriter.WriteString(outputLine + "\n")
		if err != nil {
			panic(err)
		}
	}

	// Flush slik at all data er skrvet i filen
	err = outputWriter.Flush()
	if err != nil {
		log.Fatal(err)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Ferdig!")
}
func convertLastField(lastField string) (string, error) {
	if lastField == "" {
		return "", fmt.Errorf("last field is empty")
	}
	// Konverterer siste field til en float
	celsius, err := strconv.ParseFloat(lastField, 64)
	if err != nil {
		return "", err
	}
	fahrenheit := conv.CelsiusToFahrenheit(celsius)
	return fmt.Sprintf("%.1f", fahrenheit), nil
}

func AverageTemperature() {

	// Åpner csv filen
	file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Leser csv filen
	reader := csv.NewReader(file)
	var lines [][]string
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, line)
	}

	fmt.Println("Velg temperaturenhet (celsius/fahr):")
	var unit string
	fmt.Scan(&unit)

	// Kalkulerer gjennomsnitlig temperatur
	var sum float64
	count := 0

	for i, fields := range lines {
		if i == 0 {
			continue
		}
		if len(fields) != 4 {
			log.Fatalf("unexpected number of fields in line %d: %d", i, len(fields))
		}
		if fields[3] == "" {
			continue
		}
		temperature, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			log.Fatalf("could not parse temperature in line %d: %s", i, err)
		}
		if unit == "fahr" {
			temperature = conv.CelsiusToFahrenheit(temperature)
		}
		sum += temperature
		count++
	}
	if unit == "fahr" {
		average := sum / float64(count)
		average = math.Round(average*100) / 100

		fmt.Printf("Gjennomsnittlig temperatur: %.2f°F\n", average)
	} else {
		average := sum / float64(count)
		fmt.Printf("Gjennomsnittlig temperatur: %.2f°C\n", average)
	}

}
