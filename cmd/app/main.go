package main

import (
	"github.com/FuVy/GSheetsToUnicode"
	"github.com/FuVy/GSheetsToUnicode/pkg/env"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	// Your Google API key
	apiKey := env.GetOrPanicOnEmpty("API_KEY")

	// The ID of the spreadsheet to retrieve data from.
	spreadsheetId := env.GetOrPanicOnEmpty("SPREADSHEET_ID")
	//readRange := env.GetOrPanicOnEmpty("READ_RANGE")

	inputPath := env.GetOrPanicOnEmpty("INPUT_FILE")
	outputPath := env.GetOrPanicOnEmpty("OUTPUT_FILE")
	GSheetsToUnicode.Parse(apiKey, spreadsheetId, inputPath, outputPath)
}
