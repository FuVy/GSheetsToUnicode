package GSheetsToUnicode

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func Parse(apiKey, spreadsheetId, inputPath string) (symbols []string, err error) {
	ctx := context.Background()
	characters := make(map[string]interface{})

	lines, err := readLines(inputPath) // Replace with your file path
	if err != nil {
		log.Println("Error reading file:", err)
		return
	}

	srv, err := sheets.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Printf("Unable to retrieve Sheets client: %v\n", err)
		return
	}
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, readRange := range lines {
		wg.Add(1)
		go func(readRange string) {
			defer wg.Done()
			// Call the Sheets API to get the values
			resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
			if err != nil {
				log.Printf("Unable to retrieve data from sheet: %v\n", err)
				return
			}

			if len(resp.Values) == 0 {
				log.Println("No data found.")
				return
			} else {
				for _, row := range resp.Values {
					str := fmt.Sprintf("%v", row)
					addUnicodeHexCodes(str, &characters, &mu)
				}
			}
			fmt.Printf("%v completed \n", readRange)
		}(readRange)
	}
	wg.Wait()

	slice := valuesToStringSlice(characters)

	return slice, nil
}

func ParseAndWrite(apiKey, spreadsheetId, inputPath, outputPath string) (err error) {
	slice, err := Parse(apiKey, spreadsheetId, inputPath)
	if err != nil {
		log.Printf("error parsing sheets: %v\n", err)
		return
	}
	err = writeToFile(outputPath, slice)
	if err != nil {
		log.Printf("error writing results to file: %v\n", err)
		return
	}
	return nil
}

// readLines reads a file and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func addUnicodeHexCodes(input string, unicodeHexCodes *map[string]interface{}, mu *sync.Mutex) {
	for _, r := range input {
		hexCode := fmt.Sprintf("%04X", r) // Convert rune to 4-digit hexadecimal string
		mu.Lock()
		(*unicodeHexCodes)[hexCode] = true
		mu.Unlock()
	}
}

func writeToFile(filename string, data []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(strings.Join(data, ","))
	if err != nil {
		return err
	}
	return nil
}

func valuesToStringSlice(m map[string]interface{}) []string {
	array := make([]string, 0, len(m))
	for k := range m {
		array = append(array, k)
	}
	return array
}
