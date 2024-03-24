package main

import (
	"github.com/FuVy/GSheetsToUnicode"
	"github.com/FuVy/GSheetsToUnicode/pkg/env"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	apiKey := env.GetOrPanicOnEmpty("API_KEY")
	spreadsheetId := env.GetOrPanicOnEmpty("SPREADSHEET_ID")
	inputPath := env.GetOrPanicOnEmpty("INPUT_FILE")
	outputPath := env.GetOrPanicOnEmpty("OUTPUT_FILE")

	GSheetsToUnicode.ParseAndWrite(apiKey, spreadsheetId, inputPath, outputPath)
}
