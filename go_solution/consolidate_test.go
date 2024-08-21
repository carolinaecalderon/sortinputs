package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConsolidateData(t *testing.T) {
	cases := []string{"simple", "chaos", "complex"}

	for _, tt := range cases {
		t.Run(tt, func(t *testing.T) {
			input_directory := "../tests/" + tt
			output_expected := input_directory + ".txt"
			output_actual := input_directory + "-output.txt"

			err := consolidateData(input_directory, output_actual)
			require.NoError(t, err)

			// Read the contents of the files
			expected, err := os.ReadFile(output_expected)
			if err != nil {
				t.Fatalf("Failed to read %s: %v", output_expected, err)
			}

			actual, err := os.ReadFile(output_actual)
			if err != nil {
				t.Fatalf("Failed to read %s: %v", output_actual, err)
			}

			// Compare the contents
			assert.Equal(t, expected, actual)

			// Delete created files
			if err = os.Remove(output_actual); err != nil {
				t.Fatalf("Failed to clean up output file %s: %v", output_actual, err)
			}

		})
	}
}
