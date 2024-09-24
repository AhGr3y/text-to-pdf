package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	entries, err := os.ReadDir("./text_files")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, entry := range entries {
		pdf := gofpdf.New("Portrait", "mm", "A4", "")
		chapterBody := func(fileStr string) {
			// Read text file
			fmt.Printf("Reading %s...\n", fileStr)
			dat, err := os.ReadFile(fileStr)
			if err != nil {
				pdf.SetError(err)
			}
			// Times 12
			pdf.SetFont("Times", "", 12)
			// Output justified text
			pdf.MultiCell(0, 5, string(dat), "", "", false)
			// Line break
			pdf.Ln(-1)
		}
		printText := func(fileStr string) {
			pdf.AddPage()
			chapterBody(fileStr)
		}
		fmt.Printf("Found: %v\n", entry.Name())
		if entry.IsDir() {
			fmt.Println("Skipping dir...")
			continue
		} else {
			printText("./text_files/" + entry.Name())
			fileName, hasPrefix := strings.CutSuffix(entry.Name(), ".txt")
			if !hasPrefix {
				fmt.Printf("Not a text file, skipping it: %s", entry.Name())
				continue
			}
			fileStr := fileName + ".pdf"
			err := pdf.OutputFileAndClose("./pdf_files/" + fileStr)
			if err != nil {
				pdf.SetError(err)
			}
			fmt.Printf("Done with %v...\n", entry.Name())
		}
	}
}
