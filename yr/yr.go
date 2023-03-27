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

// Konvertere temperatur
func ConvertTemperature() {
	overwriteFile := checkFileExists()
	if !overwriteFile {
		fmt.Println("Going back to main menu")
		return
	}

	inputFile := openInputFile()
	defer inputFile.Close()

	outputFile, err := createOutputFile()
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)

	scanner := bufio.NewScanner(inputFile)

	if scanner.Scan() {
		_, err := outputWriter.WriteString(scanner.Text() + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	for scanner.Scan() {
		line := scanner.Text()

		// Prosesser input-linje
		outputLine := ProcessLine(line)
		if err != nil {
			log.Fatalf("error processing input line: %v", err)
		}

		// Skriv ferdig prosessert input linje til output-fil
		_, err = outputWriter.WriteString(outputLine + "\n")
		if err != nil {
			log.Fatalf("error writing to output file: %v", err)
		}
	}

	err = outputWriter.Flush()
	if err != nil {
		log.Fatalf("error flushing output writer: %v", err)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error scanning input file: %v", err)
	}

	fmt.Println("Completed!")
}

// Sjekker om filen allerede eksisterer
func checkFileExists() bool {
	if _, err := os.Stat("kjevik-temp-fahr-20220318-20230318.csv"); err == nil {
		fmt.Printf("The file already exists. Do you want to overwrite it? (y/n): ")

		var overwriteInput string
		fmt.Scanln(&overwriteInput)

		if strings.ToLower(overwriteInput) == "y" {
			err := os.Remove("kjevik-temp-fahr-20220318-20230318.csv")
			if err != nil {
				log.Fatal(err)
			}
			return true
		}
		return false
	}
	return true
}

// Åpner opp cvs filen
func openInputFile() *os.File {
	file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

// Lager outputfilen/cvs filen
func createOutputFile() (*os.File, error) {
	outputFilePath := "kjevik-temp-fahr-20220318-20230318.csv"
	if _, err := os.Stat(outputFilePath); err == nil {
		fmt.Printf("File %s already exists. Deleting...\n", outputFilePath)
		err := os.Remove(outputFilePath)
		if err != nil {
			return nil, fmt.Errorf("could not delete file: %v", err)
		}
	}
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not create file: %v", err)
	}
	return outputFile, nil
}

// Prosesserer input data
func ProcessLine(line string) string {
	if line == "" {
		return ""
	}
	fields := strings.Split(line, ";")
	lastField := ""
	if len(fields) > 0 {
		lastField = fields[len(fields)-1]
	}
	if lastField != "" {
		var err error
		lastField, err = convertLastField(lastField)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return ""
		}
		fields[len(fields)-1] = lastField
	}
	if line[0:7] == "Data er" {
		return "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Thomas"
	} else {
		return strings.Join(fields, ";")
	}
}

func convertLastField(lastField string) (string, error) {
	// Konverterer nummer i string til float64
	celsius, err := strconv.ParseFloat(lastField, 64)
	if err != nil {
		return "", err
	}

	// Konverterer Celsius til Farenheit
	fahrenheit := (celsius * 9 / 5) + 32

	// Konverterer float64 Fahrenheit tilbake til string
	return strconv.FormatFloat(fahrenheit, 'f', 1, 64), nil
}

func AverageTemperature() {

	// Åpne input-fil
	file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Opprett en csv.Reader for å lese fila
	reader := csv.NewReader(file) // Bruker csv.reader
	reader.Comma = ';'

	// Les og ignorer header-linjen
	_, err = reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	// Be brukeren om å skrive gjennomsnittlig temperatur i celsius eller fahrenheit
	fmt.Println("Choose unit of temperature (c or f):")
	var unit string
	fmt.Scan(&unit)

	// Regne ut gjennomsnittlig temperatur
	var sum float64
	count := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(record) != 4 {
			log.Fatalf("unexpected number of fields in line %v", record)
		}
		if record[3] == "" {
			continue
		}
		temperature, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			log.Fatalf("could not parse temperature in line %v: %s", record, err)
		}
		if unit == "f" {
			// Konverterer tilbake til fahrenheit om det var det brukeren skrev inn
			temperature = conv.CelsiusToFahrenheit(temperature)
		}
		sum += temperature
		count++
	}

	if unit == "f" {
		average := sum / float64(count)
		average = math.Round(average*100) / 100
		fmt.Printf("Average temperature in Farenheit: %.2f°F\n", average)
	} else {
		average := sum / float64(count)
		fmt.Printf("Average temperature in Celsius: %.2f°C\n", average)
	}
}

// Funksjon som teller linjer i en fil

func CountLines(inputFile string) (int, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	countedLines := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			countedLines++
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return countedLines, nil
}

// Tar inn input og retunerer gjennomsnittstemperatur basert på input
func GetAverageTemperature(filepath string, unit string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var sum float64
	count := 0
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		if i == 0 {
			continue
		}
		fields := strings.Split(scanner.Text(), ";")
		if len(fields) != 4 {
			return "", fmt.Errorf("unexpected number of fields in line %d: %d", i, len(fields))
		}
		if fields[3] == "" {
			continue
		}
		temperature, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			return "", fmt.Errorf("could not parse temperature in line %d: %s", i, err)
		}

		if unit == "f" {
			temperature = conv.CelsiusToFahrenheit(temperature)
		}
		sum += temperature
		count++
	}
	average := sum / float64(count)
	return fmt.Sprintf("%.2f", average), nil

}
