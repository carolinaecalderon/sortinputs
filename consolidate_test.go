package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConsolidateData(t *testing.T) {
	cases := []struct {
		name         string
		inputDir     string
		outputFile   string
		outputActual string
	}{
		{"simple", "tests/simple", "tests/simple-output.txt", "tests/simple.txt"},
		{"complex", "tests/complex", "tests/complex-output.txt", "tests/complex.txt"},
		{"chaos", "tests/chaos", "tests/chaos-output.txt", "tests/chaos.txt"},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := consolidateData(tt.inputDir, tt.outputFile)
			require.NoError(t, err)

			// Read the contents of the files
			expected, err := os.ReadFile(tt.outputFile)
			if err != nil {
				t.Fatalf("Failed to read file1.txt: %v", err)
			}

			actual, err := os.ReadFile(tt.outputActual)
			if err != nil {
				t.Fatalf("Failed to read file2.txt: %v", err)
			}

			// Compare the contents
			assert.Equal(t, expected, actual)
		})
	}
}
