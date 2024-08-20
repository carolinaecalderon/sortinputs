package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

// consolidateData calls the goroutine processFile function on each file inside the
// input directory to parse out the non-blank text lines into a string channel
// then, it sorts the unique text lines into one output file
func consolidateData(inputDir string, outputFile string) error {
	// initialize variables
	lines := make(chan string, 0)
	uniqueLines := map[string]bool{}
	sortedLines := []string{}

	// check that the inputDir is not empty/can be walked
	info, err := os.Stat(inputDir)
	if err != nil {
		return fmt.Errorf("unable to recognize input directory %s: %s", inputDir, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("input directory %s is not a directory", inputDir)
	}

	// collect all the lines
	var wg sync.WaitGroup
	if err := filepath.WalkDir(inputDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			wg.Add(1)
			go func() {
				defer wg.Done()
				processFile(path, lines)
			}()
		}
		return nil
	}); err != nil {
		return err
	}

	// wait until all files are walked & close the string channel
	go func() {
		wg.Wait()
		close(lines)
	}()

	// now only consider unique lines
	for line := range lines {
		// if this doesn't exist in the unique map already, we know it's our first time coming across this data
		if !uniqueLines[line] {
			// add it to the map
			uniqueLines[line] = true
			// add it to an array (to later sort)
			sortedLines = append(sortedLines, line)
		}
	}

	// sort the unique lines into a string list/array
	sort.Strings(sortedLines)

	// write the output file & return error/nil
	return writeFile(outputFile, sortedLines)
}

// processFile parses out the non-blank text lines from a given file path
func processFile(fp string, lines chan<- string) {
	if filepath.Ext(fp) != ".txt" {
		return // Skip files that aren't .txt
	}

	// try to open the file
	file, err := os.Open(fp)
	if err != nil {
		fmt.Printf("error opening file: %s", fp)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines <- line
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("error reading file: %s", fp)
	}
	return
}

// writeFile writes the sorted lines to the output file
func writeFile(fp string, sortedLines []string) error {
	file, err := os.Create(fp)
	if err != nil {
		return fmt.Errorf("error creating output file %s", fp)
	}
	defer file.Close()

	for _, v := range sortedLines {
		if _, err := file.WriteString(v + "\n"); err != nil {
			return err
		}
	}
	return nil
}
