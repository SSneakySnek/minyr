package yr_test

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/SSneakySnek/minyr/yr"
)

// Tester antall linjer i filen
func TestFileLineCount(t *testing.T) {
	tests := []struct {
		filename         string
		expectedNumLines int
	}{
		{"kjevik-temp-celsius-20220318-20230318.csv", 16756},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("count lines in %s", tc.filename), func(t *testing.T) {
			file, err := os.Open(tc.filename)
			if err != nil {
				t.Fatalf("Failed to open file %s: %v", tc.filename, err)
			}
			defer file.Close()

			var numLines int
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				numLines++
			}
			if err := scanner.Err(); err != nil {
				t.Fatalf("Failed to scan file %s: %v", tc.filename, err)
			}

			if numLines != tc.expectedNumLines {
				t.Errorf("Expected %v lines in %v, but got %v lines", tc.expectedNumLines, tc.filename, numLines)
			}
		})
	}
}

// Tester på å konvertere linjer
func TestConvertLines(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "valid input", input: "Kjevik;SN39040;18.03.2022 01:50;6", want: "Kjevik;SN39040;18.03.2022 01:50;42.8"},
		{name: "valid input with zero temperature", input: "Kjevik;SN39040;07.03.2023 18:20;0", want: "Kjevik;SN39040;07.03.2023 18:20;32.0"},
		{name: "valid input with negative temperature", input: "Kjevik;SN39040;08.03.2023 02:20;-11", want: "Kjevik;SN39040;08.03.2023 02:20;12.2"},
		{name: "invalid input", input: "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;", want: "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Thomas"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := yr.ProcessLine(tt.input); got != tt.want {
				t.Errorf("ProcessLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Tester om gjennomsnittstemeratur er 8.56
func TestGetAverageTemperature(t *testing.T) {
	actualAvg, err := yr.GetAverageTemperature("kjevik-temp-celsius-20220318-20230318.csv", "celsius")
	if err != nil {
		t.Fatal(err)
	}

	expectedAvg := "8.56"
	if actualAvg != expectedAvg {
		t.Errorf("expected average temperature %v, but got %v", expectedAvg, actualAvg)
	}
}
