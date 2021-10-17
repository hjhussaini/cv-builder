package models

type Project struct {
	Name           string   `yaml:"name"`
	StartDate      string   `yaml:"start_date"`
	EndDate        string   `yaml:"end_date"`
	Creator        string   `yaml:"creator"`
	AssociatedWith string   `yaml:"associated_with"`
	URL            string   `yaml:"url"`
	Description    []string `yaml:"description"`
}
