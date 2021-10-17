package main

import (
	"fmt"
	"os"

	"github.com/hjhussaini/cv-builder/models"
	"github.com/hjhussaini/cv-builder/resume"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <YAML FILE> <PDF FILE NAME>\n", os.Args[0])
		os.Exit(1)
	}
	yamlFile := os.Args[1]
	fileName := os.Args[2]

	sections := &models.Sections{}
	if err := models.ReadYAMLFile(yamlFile, sections); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := resume.BuildPDF(sections, fileName); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
