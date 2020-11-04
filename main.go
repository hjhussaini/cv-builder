package main

import (
	"fmt"
	"os"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <PDF FILE NAME>\n", os.Args[0])
		os.Exit(1)
	}
	fileName := os.Args[1]

	pdf := gofpdf.New("Portrait", "cm", "A4", "")
	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
