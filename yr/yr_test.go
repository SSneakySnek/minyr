package yr_test

import (
	"bufio"
	"os"
	"testing"

	"github.com/SSneakySnek/minyr/yr"
)

func TestFileLineCount(t *testing.T) {
	filename := "kjevik-temp-celsius-20220318-20230318.csv"
	expectedNumLines := 16756

	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	var numLines int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numLines++
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("Failed to scan file: %v", err)
	}

	if numLines != expectedNumLines {
		t.Errorf("Expected %v lines in %v, but got %v lines", expectedNumLines, filename, numLines)
	}
}

func TestConvertLines(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "Kjevik;SN39040;18.03.2022 01:50;6", want: "Kjevik;SN39040;18.03.2022 01:50;42.8"},
		{input: "Kjevik;SN39040;07.03.2023 18:20;0", want: "Kjevik;SN39040;07.03.2023 18:20;32.0"},
		{input: "Kjevik;SN39040;08.03.2023 02:20;-11", want: "Kjevik;SN39040;08.03.2023 02:20;12.2"},
		{input: "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;",
			want: "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Thomas"},
	}

	for _, tc := range tests {
		got := yr.ProcessLine(tc.input)
		if tc.want != got {
			t.Errorf("Expected %v, but got %v", tc.want, got)
		}
	}
}
