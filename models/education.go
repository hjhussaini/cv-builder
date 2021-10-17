package models

import "fmt"

type Education struct {
	School    string `yaml:"school"`
	Degree    string `yaml:"degree"`
	Field     string `yaml:"field"`
	StartYear uint16 `yaml:"start_year"`
	EndYear   uint16 `yaml:"end_year"`
}

func (education *Education) FieldOfStudy() string {
	return education.Degree + " Degree, " + education.Field
}

func (education *Education) Date() string {
	return fmt.Sprintf("%d - %d", education.StartYear, education.EndYear)
}
