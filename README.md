# GSheetsToUnicode
 
This Go package provides a simple way to extract Unicode characters from a Google Sheets document and write them to a file.

## Installation
To use this package, you need to have Go installed on your machine. Then, you can install the package by running:

```bash
go get github.com/FuVy/GSheetsToUnicode
```
## Usage
To use this package, you need to provide the following environment variables:

- **API_KEY**: Your Google API key.
- **SPREADSHEET_ID**: The ID of the Google Sheets document.
- **INPUT_FILE**: The path to the input file containing the ranges to read from the Google Sheets document.
- **OUTPUT_FILE**: The path to the output file where the Unicode characters will be written.

You can set these environment variables in your shell or in a .env file.

Here's an example of how to use this package:

```go
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
	inputPath := env.GetOrPanicOnEmpty("INPUT_FILE")
	outputPath := env.GetOrPanicOnEmpty("OUTPUT_FILE")

	GSheetsToUnicode.ParseAndWrite(apiKey, spreadsheetId, inputPath, outputPath)
}
```
